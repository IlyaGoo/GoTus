package web_server

import (
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

func (p *WebServer) StartWebServer() {
	router := mux.NewRouter()

	router.Handle("/", http.FileServer(http.Dir("./views/")))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	logrus.Fatal(http.ListenAndServe(":"+viper.GetString("port"), router))
}
