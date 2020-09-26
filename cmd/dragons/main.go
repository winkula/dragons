package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/winkula/dragons/pkg/model"
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
		fmt.Printf("Executed in %s", elapsed)
	} else {
		help()
		os.Exit(1)
	}
}

func printStats(puzzle *model.Grid, solution *model.Grid) {
	if puzzle != nil {
		undef := puzzle.CountSquares(model.SquareUndefined)
		all := puzzle.Size()
		fmt.Printf(" > Undefinedness: %.2f (%v/%v)\n", (100.0 * float64(undef) / float64(all)), undef, all)

		difficulty := model.GetDifficulty(puzzle).String()
		fmt.Println(" > Difficulty:", difficulty)
	}

	if solution != nil {
		fmt.Println(" > Interestingness:", solution.Interestingness())
	}
}

func seed() {
	// seed the random generator
	rand.Seed(time.Now().UnixNano())
}
