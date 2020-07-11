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
		g := parse(args[1], true)
		enumerate(g)
		return
	}

	if cmd == "gen" {
		g := parse(args[1], false)
		generate(g)
		return
	}

	if cmd == "genr" {
		g := parse(args[1], false)
		generateRandom(g.Width, g.Height)
		return
	}

	if cmd == "solve" {
		g := parse(args[1], true)
		solve(g)
		return
	}

	if cmd == "val" {
		g := parse(args[1], true)
		validate(g)
		return
	}
}

func parse(s string, print bool) *model.Grid {
	g := model.Parse(s)
	if print {
		fmt.Println("Grid:")
		fmt.Println(g)
	}
	return g
}

func enumerate(g *model.Grid) {
	sucs := model.Enumerate(g)
	sort.Slice(sucs, func(i, j int) bool {
		return sucs[i].Interestingness() < sucs[j].Interestingness()
	})
	for _, s := range sucs {
		fmt.Println(s)
		fmt.Println("Interestingness:", s.Interestingness())
	}
	fmt.Printf("%v possible grids found.\n", len(sucs))
}

func generate(g *model.Grid) {
	if g.Size() == g.CountSquares(model.SquareUndefined) {
		g = model.MostInteresting(g)
	}
	fmt.Println("Grid:")
	fmt.Println(g)

	gen := model.GenerateFrom(g, model.DifficultyHard)
	fmt.Println("Generated:")
	fmt.Println(gen)
	checkDifficulty(gen)
}

func generateRandom(width int, height int) {
	g := model.Generate(width, height, model.DifficultyEasy)
	fmt.Println("Generated:")
	fmt.Println(g)
	checkDifficulty(g)
}

func checkDifficulty(g *model.Grid) {
	fmt.Println("Number of undefined squares:", g.CountSquares(model.SquareUndefined))
	if model.Solve(g) != nil {
		fmt.Println("Difficulty: EASY")
	} else {
		fmt.Println("Difficulty: MEDIUM/HARD")
	}
}

func solve(g *model.Grid) {
	type solver func(g *model.Grid) *model.Grid

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
		solved := s.solver(g)
		if solved == nil {
			fmt.Println("   No solution found. Reasons can be: the puzzle is too difficult, puzzle has no distinct solution...")
			continue
		}
		fmt.Println("Solved:")
		fmt.Println(solved)
		break
	}
}

func validate(g *model.Grid) {
	valid := model.Validate(g)
	fmt.Println("Valid:", valid)
}

func seed() {
	// seed the random generator
	rand.Seed(time.Now().UnixNano())
}
