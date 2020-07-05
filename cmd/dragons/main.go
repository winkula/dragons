package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/winkula/dragons/pkg/generator"
	"github.com/winkula/dragons/pkg/model"
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

	if cmd == "gen" {
		width, _ := strconv.ParseInt(args[1], 10, 32)
		height, _ := strconv.ParseInt(args[2], 10, 32)
		generate(int(width), int(height))
		return
	}
}

func parse(s string) *model.World {
	world := model.ParseWorld(s)
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

func generate(width int, height int) {
	world := generator.GenerateWorld(width, height)
	fmt.Println("World:")
	fmt.Println(world)
	obf := generator.ObfuscateWorld(world)
	fmt.Println("Obfuscated:")
	fmt.Println(obf)
}
