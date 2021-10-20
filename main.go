package main

import (
	"GoTus/web_presenter"
	"encoding/json"
	"log"
	"os"
)

type Configuration struct {
	Version string
	Port    string
}

func main() {
	testUser := User{0, "IlyaGo", "kek"}
	testUser.Nickname = "Lune"
	run()
}

func run() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	configFile, _ := os.Open("config.json")
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		errorLog.Fatal(err)
	}

	testWebPresenter := web_presenter.WebPresenter{conf.Port, ""}
	infoLog.Printf("Web server starting on port: ", testWebPresenter.Port)
	testWebPresenter.StartWebPresenter()
}
