package main

import (
	"fmt"

	lib "github.com/ns-pkgs/golang-lib-private"
	_ "github.com/twharmon/gouid"
)

func main() {
	fmt.Println("Hello World")
	callPrint()
}

func callPrint() {
	lib.Printer()
}
