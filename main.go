package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Version string
	Port    string
}

func main() {
	testUser := User{0, "IlyaGo", "kek"}
	testUser.nickname = "Lune"
	run()
}

func run() {
	configFile, _ := os.Open("config.json")
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	conf := Configuration{}
	err := decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error:", err)
	}

	testWebPresenter := webPresenter{conf}
	testWebPresenter.StartWebPresenter()
}
