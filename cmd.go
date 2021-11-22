package cmd

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

// Stores known subcommands.
type Map map[string]func()

// Adds a command, panicking if the key is redundant or if the function is nil.
func (self Map) Add(key string, val func()) {
	if val == nil {
		panic(fmt.Errorf(`[cmd] nil command %q`, key))
	}
	if self[key] != nil {
		panic(fmt.Errorf(`[cmd] redundant command %q`, key))
	}
	self[key] = val
}

/*
Uses `Args` to find the next subcommand. Modifies `os.Args` to adjust the
remaining args for subsequent `flag.Parse` calls.
*/
func (self Map) Get() func() {
	args := Args()
	if len(args) == 0 {
		panic(fmt.Errorf(
			`[cmd] missing command; known commands: %q`,
			self.Keys(),
		))
	}

	cmd := args[0]
	fun, ok := self[cmd]
	if !ok {
		panic(fmt.Errorf(
			`[cmd] unrecognized command %q; known commands: %q`,
			cmd, self.Keys(),
		))
	}

	os.Args = args
	return fun
}

// Keys of known subcommands, sorted alphabetically.
func (self Map) Keys() []string {
	out := make([]string, 0, len(self))
	for key := range self {
		out = append(out, key)
	}
	sort.Strings(out)
	return out
}

/*
Returns remaining args from either `flag.Args` or `os.Args`, depending if
`flag.Parse` has been called.
*/
func Args() []string {
	if flag.Parsed() {
		return flag.Args()
	}
	return os.Args[1:]
}

/*
Must be deferred. Tool for panic recovery and logging. If there was a non-nil
panic, prints details to flag output (standard error by default) and exits with
a non-zero status. Otherwise, it's a nop.
*/
func Report() {
	val := recover()
	if val != nil {
		fmt.Fprintf(flag.CommandLine.Output(), "%+v\n", val)
		os.Exit(1)
	}
}
