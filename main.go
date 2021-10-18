package main

import (
	"fmt"
	"strconv"
)

func printControlResult(control Control, needPrefix bool) {
	if needPrefix {
		msg := "Контроль объекта с id = " + strconv.Itoa(control.objectId) + " : " + control.result
		fmt.Println(msg)
	} else {
		fmt.Println(control.result)
	}
}

func main() {
	testControl := Control{0, "Accepted"}
	testControl.result = "Rejected"
	printControlResult(testControl, true)
}
