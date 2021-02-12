package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

func init() {
	cmd := flag.NewFlagSet("solve", flag.ExitOnError)
	dificulty := cmd.String("difficulty", "hard", "difficulty of the puzzle")
	algorithm := cmd.String("algorithm", "all", "algorithm to use (possible values: t, i, dk)")

	registerCommand("solve", cmd, func() {
		difficultyEnum := model.ParseDifficulty(*dificulty)
		g := parse(cmd.Arg(0), true)
		solve(g, difficultyEnum, *algorithm)
	})
}

func solve(g *model.Grid, difficulty model.Difficulty, algorithm string) {
	type solver func(g *model.Grid) *model.Grid

	solvers := []struct {
		name   string
		key    string
		solver solver
	}{
		{name: "SolveDomainKnowledge", key: "t", solver: func(g *model.Grid) *model.Grid { return model.SolveTechnically(g, difficulty) }},
		{name: "SolveIterative", key: "i", solver: func(g *model.Grid) *model.Grid { return model.SolveIterative(g, difficulty) }},
		{name: "SolveBruteForce", key: "bf", solver: model.SolveBruteForce},
	}

	for _, s := range solvers {
		if algorithm == "all" || s.key == algorithm {
			fmt.Println("-> Using solver algorithm:", s.name)
			solution := s.solver(g)
			if solution == nil {
				fmt.Println("   No solution found. Reasons can be: the puzzle is too difficult, puzzle has no distinct solution...")
				continue
			} else {
				fmt.Println("Solution:")
				fmt.Println(solution)
				return
			}
		}
	}

	fmt.Println("No solver algorithm found with name ", algorithm)
}
