package main

import (
	"fmt"
	"os"

	"github.com/winkula/dragons/pkg/model"
	"github.com/winkula/dragons/pkg/parser"
)

func main() {
	args := os.Args[1:]

	cmd := args[0]

	if cmd == "parse" {
		arg := args[1]
		parse(arg)
		return
	}

	if cmd == "enum" {
		arg := args[1]
		world := parse(arg)
		enumerate(world)
		return
	}
}

func parse(s string) *model.World {
	world := parser.ParseWorld(s)
	fmt.Println("World:")
	fmt.Println(world)
	fmt.Printf("Valid: %t\n", model.ValidateWorld(world))
	return world
}

func enumerate(world *model.World) {
	successors := world.Enumerate()
	fmt.Printf("Successors (%v):\n", len(successors))
	for _, s := range successors {
		fmt.Println(s)
	}
}
