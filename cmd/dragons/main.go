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
		parse(args[1], true)
		return
	}

	if cmd == "enum" {
		world := parse(args[1], true)
		enumerate(world)
		return
	}

	if cmd == "gen" {
		world := parse(args[1], false)
		generate(world)
		return
	}

	if cmd == "genr" {
		world := parse(args[1], false)
		generateRandom(world.Width, world.Height)
		return
	}

	if cmd == "solve" {
		world := parse(args[1], true)
		solve(world)
		return
	}

	if cmd == "val" {
		world := parse(args[1], true)
		validate(world)
		return
	}
}

func parse(s string, print bool) *model.World {
	world := model.Parse(s)
	if print {
		fmt.Println("World:")
		fmt.Println(world)
	}
	return world
}

func enumerate(world *model.World) {
	successors := model.Enumerate(world)
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

	g := model.GenerateFrom(world, model.DifficultyEasy)
	fmt.Println("Generated:")
	fmt.Println(g)
	checkDifficulty(g)
}

func generateRandom(width int, height int) {
	g := model.Generate(width, height, model.DifficultyEasy)
	fmt.Println("Generated:")
	fmt.Println(g)
	checkDifficulty(g)
}

func checkDifficulty(w *model.World) {
	fmt.Println("Number of undefined squares:", w.CountSquares(model.SquareUndefined))
	if model.Solve(w) != nil {
		fmt.Println("Difficulty: EASY")
	} else {
		fmt.Println("Difficulty: MEDIUM/HARD")
	}
}

func solve(world *model.World) {
	type solver func(w *model.World) *model.World

	solvers := []struct {
		name   string
		solver solver
	}{
		{name: "SolveDk", solver: model.SolveDk},
		{name: "SolveBf", solver: model.SolveBf},
		{name: "Solve", solver: model.Solve},
	}

	for _, s := range solvers {
		fmt.Println("-> Using solver algorithm:", s.name)
		solved := s.solver(world)
		if solved == nil {
			fmt.Println("   No solution found. Reasons can be: the puzzle is too difficult, puzzle has no distinct solution...")
			continue
		}
		fmt.Println("Solved:")
		fmt.Println(solved)
		break
	}
}

func validate(world *model.World) {
	valid := model.Validate(world)
	fmt.Println("IsValid:", valid)
}

func seed() {
	// seed the random generator
	rand.Seed(time.Now().UnixNano())
}
