package main

import (
	"flag"
	"fmt"

	"github.com/mitranim/cmd"
)

var (
	commands = cmd.Map{}
	flagA    = flag.Bool(`a`, false, `flag "a"`)
)

func main() {
	flag.Parse()
	fmt.Printf(`running with "-a" = %v`+"\n", *flagA)

	defer cmd.Report()
	commands.Get()()
}

func init() { commands.Add(`one`, cmdOne) }

func cmdOne() {
	flagB := flag.Bool(`b`, false, `flag "b"`)

	flag.Parse()

	fmt.Printf(
		`running command %q with "-b" = %v and args = %q`+"\n",
		`one`, *flagB, cmd.Args(),
	)
}

func init() { commands.Add(`two`, cmdTwo) }

func cmdTwo() {
	flagC := flag.Bool(`c`, false, `flag "c"`)

	flag.Parse()

	fmt.Printf(
		`running command %q with "-c" = %v and args = %q`+"\n",
		`two`, *flagC, cmd.Args(),
	)
}
