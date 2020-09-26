package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

var solveCmd = flag.NewFlagSet("solve", flag.ExitOnError)
var solArgDifficulty = solveCmd.String("difficulty", "easy", "difficulty of the puzzle")

func init() {
	registerCommand("solve", solveCmd, func() {
		difficultyEnum := model.ParseDifficulty(*solArgDifficulty)
		g := parse(solveCmd.Arg(0), true)
		solve(g, difficultyEnum)
	})
}

func solve(g *model.Grid, difficulty model.Difficulty) {
	type solver func(g *model.Grid) (*model.Grid, *model.SolveResult)

	solvers := []struct {
		name   string
		solver solver
	}{
		{name: "Solve", solver: func(g *model.Grid) (*model.Grid, *model.SolveResult) {
			return model.SolveHuman(g, difficulty)
		}},
		{name: "SolveDk", solver: model.SolveDk},
		{name: "SolveBf", solver: model.SolveBf},
	}

	for _, s := range solvers {
		fmt.Println("-> Using solver algorithm:", s.name)
		solved, solveResult := s.solver(g)
		if solved == nil || solveResult.Difficulty > difficulty {
			fmt.Println("   No solution found. Reasons can be: the puzzle is too difficult, puzzle has no distinct solution...")
			continue
		}
		fmt.Println("Solved:")
		fmt.Println(solved)
		if solveResult != nil {
			fmt.Println("> Permutations:", solveResult.MaxPerm)
		}
		break
	}
}
