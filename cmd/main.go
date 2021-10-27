package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	//Mat Ryer advice to handle all app errors
	if err := execute(); err != nil {
		logrus.Fatal(err)
	}
}
