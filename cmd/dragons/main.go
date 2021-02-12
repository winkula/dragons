package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var commands = map[string]command{}

type command struct {
	flags  *flag.FlagSet
	action func()
}

func registerCommand(name string, flags *flag.FlagSet, action func()) {
	commands[name] = command{
		flags:  flags,
		action: action,
	}
}

func main() {
	seed()

	if len(os.Args) < 2 {
		help()
		os.Exit(1)
		return
	}

	key := os.Args[1]
	if command, found := commands[key]; found {
		args := os.Args[2:]
		if command.flags != nil {
			command.flags.Parse(args)
		}
		start := time.Now()
		command.action()
		elapsed := time.Since(start)
		fmt.Printf("Executed in %s\n", elapsed)
	} else {
		help()
		os.Exit(1)
	}
}

func seed() {
	// seed the random generator
	rand.Seed(time.Now().UnixNano())
}
