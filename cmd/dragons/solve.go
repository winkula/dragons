package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

func init() {
	cmd := flag.NewFlagSet("solve", flag.ExitOnError)
	dificulty := cmd.String("difficulty", "easy", "difficulty of the puzzle")

	registerCommand("solve", cmd, func() {
		difficultyEnum := model.ParseDifficulty(*dificulty)
		g := parse(cmd.Arg(0), true)
		solve(g, difficultyEnum)
	})
}

func solve(g *model.Grid, difficulty model.Difficulty) {
	type solver func(g *model.Grid) *model.Grid

	solvers := []struct {
		name   string
		solver solver
	}{
		{name: "SolveHuman", solver: func(g *model.Grid) *model.Grid {
			return model.SolveHuman(g, difficulty)
		}},
		{name: "SolveDomainKnowledge", solver: model.SolveDk},
		{name: "SolveBruteForce", solver: model.SolveBf},
	}

	for _, s := range solvers {
		fmt.Println("-> Using solver algorithm:", s.name)
		solution := s.solver(g)
		if solution == nil {
			fmt.Println("   No solution found. Reasons can be: the puzzle is too difficult, puzzle has no distinct solution...")
			continue
		}
		fmt.Println("Solution:")
		fmt.Println(solution)
		break
	}
}
