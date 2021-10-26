package main

import (
	"GoTus/cmd"

	"github.com/sirupsen/logrus"
)

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
