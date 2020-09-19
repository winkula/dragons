package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/winkula/dragons/pkg/model"
)

func main() {
	seed()

	if len(os.Args) < 2 {
		help()
		os.Exit(1)
		return
	}

	start := time.Now()

	switch os.Args[1] {

	case "generate":
		generateCmd.Parse(os.Args[2:])
		difficultyEnum := model.ParseDifficulty(*genArgDifficulty)

		if len(generateCmd.Args()) > 0 {
			g := parse(generateCmd.Arg(0), false)
			generateFrom(g, difficultyEnum, *genArgDuration)
		} else {
			generate(*genArgWidth, *genArgHeight, difficultyEnum, *genArgDuration/2)
		}

	case "validate":
		validateCmd.Parse(os.Args[2:])
		g := parse(validateCmd.Arg(0), true)
		validate(g)

	case "parse":
		parseCmd.Parse(os.Args[2:])
		parse(parseCmd.Arg(0), true)

	case "enumerate":
		enumerateCmd.Parse(os.Args[2:])
		g := parse(enumerateCmd.Arg(0), true)
		enumerate(g, *enuArgMost)

	case "solve":
		solveCmd.Parse(os.Args[2:])
		difficultyEnum := model.ParseDifficulty(*solArgDifficulty)
		g := parse(solveCmd.Arg(0), true)
		solve(g, difficultyEnum)

	case "play":
		play()

	default:
		help()
		os.Exit(1)
	}

	elapsed := time.Since(start)
	fmt.Printf("Executed in %s", elapsed)
}

func printStats(puzzle *model.Grid, solution *model.Grid) {
	if puzzle != nil {
		undef := puzzle.CountSquares(model.SquareUndefined)
		all := puzzle.Size()
		fmt.Printf(" > Undefinedness: %.2f (%v/%v)\n", (100.0 * float64(undef) / float64(all)), undef, all)

		_, difficulty := getDifficulty(puzzle)
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

func getDifficulty(puzzle *model.Grid) (model.Difficulty, string) {
	solution, _ := model.SolveHuman(puzzle, model.DifficultyEasy)
	if solution != nil {
		return model.DifficultyEasy, "easy"
	}

	solution, _ = model.SolveHuman(puzzle, model.DifficultyMedium)
	if solution != nil {
		return model.DifficultyMedium, "medium"
	}

	return model.DifficultyHard, "hard"
}
