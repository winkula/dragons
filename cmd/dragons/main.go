package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/winkula/dragons/pkg/model"
)

func main() {
	seed()

	args := os.Args[1:]
	cmd := args[0]

	if cmd == "parse" {
		parse(args[1])
		return
	}

	if cmd == "enum" {
		world := parse(args[1])
		enumerate(world)
		return
	}

	if cmd == "gen" {
		world := parse(args[1])
		generate(world)
		return
	}

	if cmd == "solve" {
		world := parse(args[1])
		solve(world)
		return
	}

	if cmd == "solves" {
		world := parse(args[1])
		solveSimple(world)
		return
	}

	if cmd == "val" {
		world := parse(args[1])
		validate(world)
		return
	}
}

func parse(s string) *model.World {
	world := model.ParseWorld(s)
	fmt.Println("World:")
	fmt.Println(world)
	fmt.Println("Valid:", model.ValidateWorld(world))
	return world
}

func enumerate(world *model.World) {
	successors := world.Enumerate()
	sort.Slice(successors, func(i, j int) bool {
		return successors[i].Interestingness() < successors[j].Interestingness()
	})
	for _, s := range successors {
		fmt.Println(s)
		fmt.Println("Interestingness:", s.Interestingness())
	}
	fmt.Printf("%v possible worlds found.\n", len(successors))
}

func generate(world *model.World) {
	if world.Size() == world.CountSquares(model.SquareUndefined) {
		world = model.MostInteresting(world)
	}
	fmt.Println("World:")
	fmt.Println(world)

	g := model.GenerateWorld(world, model.DifficultyEasy)
	fmt.Println("Generated:")
	fmt.Println(g)
	if model.SolveSimple(g) != nil {
		fmt.Println("SolveSimple could solve this. Difficulty: EASY!")
	} else {
		fmt.Println("SolveSimple could not solve this. Difficulty: MEDIUM/HARD!")
	}

	fmt.Println("Undef squares:", g.CountSquares(model.SquareUndefined))
}

func solve(world *model.World) {
	solved := model.Solve(world)
	if solved == nil {
		fmt.Println("Solver: No solution found! There are two reasons:")
		fmt.Println(" - the puzzle is too difficult for the naive solver")
		fmt.Println(" - the puzzle has no distinct solution")
	} else {
		fmt.Println("Solved:")
		fmt.Println(solved)
	}
}

func solveSimple(world *model.World) {
	solved := model.SolveSimple(world)
	if solved == nil {
		fmt.Println("Solver: No solution found! There are two reasons:")
		fmt.Println(" - the puzzle is too difficult for the naive solver")
		fmt.Println(" - the puzzle has no distinct solution")
	} else {
		fmt.Println("Solved:")
		fmt.Println(solved)
	}
}

func validate(world *model.World) {
	valid := model.ValidateWorld(world)
	fmt.Println("IsValid:", valid)
}

func seed() {
	// seed the random generator
	rand.Seed(time.Now().UnixNano())
}
