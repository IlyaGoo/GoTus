package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path"

	"github.com/braintree/manners"
)

func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Привет, это учебный проект для организации тусичей")
}

func endPage(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Пока")
}

type webPresenter struct {
	config Configuration
}

func (p *webPresenter) StartWebPresenter() {
	pr := newPathResolver()
	pr.Add("GET /homePage", homePage)
	pr.Add("* /endPage/*", endPage)

	adress := "localhost:" + p.config.Port
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt, os.Kill)
	go listenForShutDown(ch)
	manners.ListenAndServe(adress, pr)
}

func (p *pathResolver) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	check := req.Method + " " + req.URL.Path

	for pattern, handlerFunc := range p.handlers {
		if ok, err := path.Match(pattern, check); ok && err == nil {
			handlerFunc(res, req)
			return
		} else if err != nil {
			fmt.Fprint(res, err)
		}
	}

	http.NotFound(res, req)
}

func listenForShutDown(ch <-chan os.Signal) {
	<-ch
	manners.Close()
}

type pathResolver struct {
	handlers map[string]http.HandlerFunc
}

func newPathResolver() *pathResolver {
	return &pathResolver{make(map[string]http.HandlerFunc)}
}

func (p *pathResolver) Add(path string, handler http.HandlerFunc) {
	p.handlers[path] = handler
}
