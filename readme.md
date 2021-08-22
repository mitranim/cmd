## Overview

Missing feature of the Go standard library: ability to define subcommands while using `flag`.

  * Complements `flag` by adding subcommands.
  * Does not reinvent flag parsing.
  * Does not pollute your stacktraces.
  * Tiny, no external dependencies.

API docs: https://pkg.go.dev/github.com/mitranim/cmd.

## TOC

* [Guideline](#guideline)
* [Usage Simple](#usage-simple)
* [Usage Advanced](#usage-advanced)

## Guideline

Store commands in a global `cmd.Map{}`. The map may be modified by `init` functions defined in different files.

Use global `flag.Parse()` and `cmd.Args()` to parse flags and obtain args. `flag.Parse()` may be called from `main` and from subcommands, multiple times, gradually consuming remaining `os.Args`.

## Usage Simple

CLI usage:

```sh
go run . --help
go run . one
go run . two
```

Go code:

```golang
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
```

## Usage Advanced

CLI usage:

```sh
go run . --help
go run . -a one -b three
go run . -a two -c three
```

Go code:

```golang
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

func init() { commands[`one`] = cmdOne }

func cmdOne() {
  flagB := flag.Bool(`b`, false, `flag "b"`)

  flag.Parse()

  fmt.Printf(
    `running command %q with "-b" = %v and args = %q`+"\n",
    `one`, *flagB, cmd.Args(),
  )
}

func init() { commands[`two`] = cmdTwo }

func cmdTwo() {
  flagC := flag.Bool(`c`, false, `flag "c"`)

  flag.Parse()

  fmt.Printf(
    `running command %q with "-c" = %v and args = %q`+"\n",
    `two`, *flagC, cmd.Args(),
  )
}
```

## License

https://unlicense.org

## Misc

I'm receptive to suggestions. If this library _almost_ satisfies you but needs changes, open an issue or chat me up. Contacts: https://mitranim.com/#contacts
