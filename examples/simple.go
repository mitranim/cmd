package main

import (
	"fmt"

	"github.com/mitranim/cmd"
)

var commands = cmd.Map{
	`one`: cmdOne,
	`two`: cmdTwo,
}

func main() {
	defer cmd.Report()
	commands.Get()()
}

func cmdOne() {
	fmt.Println(`running command "one"`)
}

func cmdTwo() {
	fmt.Println(`running command "two"`)
}
