package main

import (
	"GoTus/cmd"

	"github.com/sirupsen/logrus"
)

type Configuration struct {
	Version string
	Port    string
}

func main() {
	//Mat Ryer advice to handle all app errors
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	cmd.Execute()
	return nil
}
