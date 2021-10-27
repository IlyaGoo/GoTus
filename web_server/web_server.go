package web_server

import (
	"GoTus/administration"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/gorilla/mux"
)

type WebServer struct {
}

func NewWebServer() WebServer {
	return WebServer{}
}

var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})

var StatusHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
})

var UsersHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// Here we are converting the slice of products to JSON
	payload, _ := json.Marshal(administration.TestUsers)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(payload))
})

var AddAbboutMeHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user administration.User
	vars := mux.Vars(r)
	slug := vars["slug"]

	for _, u := range administration.TestUsers {
		if u.Slug == slug {
			user = u
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if user.Slug != "" {
		payload, _ := json.Marshal(user)
		w.Write([]byte(payload))
	} else {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}
})

func (p *WebServer) StartWebServer() {
	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./views/")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.Handle("/status", StatusHandler).Methods("GET")
	router.Handle("/users", UsersHandler).Methods("GET")
	router.Handle("/users/{slug}/abboutme", AddAbboutMeHandler).Methods("POST")

	logrus.Fatal(http.ListenAndServe(":"+viper.GetString("port"), router))
}
