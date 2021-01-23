package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/winkula/dragons/pkg/model"
	"github.com/winkula/dragons/pkg/renderers"
)

func init() {
	cmd := flag.NewFlagSet("generate", flag.ExitOnError)
	difficulty := cmd.String("difficulty", "easy", "difficulty of the puzzle")
	size := cmd.Int("size", 8, "the puzzles size")
	duration := cmd.Duration("duration", time.Second, "the available time to generate the puzzle")
	solutionOnly := cmd.Bool("solution", false, "generates only a solution")

	registerCommand("generate", cmd, func() {
		difficultyEnum := model.ParseDifficulty(*difficulty)

		if *solutionOnly {
			solution := model.Generate(*size, *size, *duration)
			fmt.Println("Solution:")
			fmt.Println(solution)
			printStats(nil, solution)
			return
		}

		if len(cmd.Args()) > 0 {
			g := parse(cmd.Arg(0), false)
			generateFrom(g, difficultyEnum, *duration)
		} else {
			generate(*size, *size, difficultyEnum, *duration/2)
		}
	})
}

func generate(width int, height int, difficulty model.Difficulty, duration time.Duration) {
	solution := model.Generate(width, height, duration)
	if solution == nil {
		fmt.Println("No puzzle could be generated with the given constraints.")
		return
	}

	generateFrom(solution, difficulty, duration)
}

func generateFrom(solution *model.Grid, difficulty model.Difficulty, duration time.Duration) {
	if solution.IsUndefined() {
		generate(solution.Width, solution.Height, difficulty, duration)
		return
	}

	fmt.Println("Solution:")
	fmt.Println(solution)

	puzzle := model.Obfuscate(solution, difficulty, duration)
	fmt.Println("Puzzle:")
	fmt.Println(puzzle)
	renderers.RenderSvg(puzzle, true)

	printStats(puzzle, solution)
}

func printStats(puzzle *model.Grid, solution *model.Grid) {
	if solution != nil {
		fmt.Println(" > Interestingness:", solution.Interestingness())
		fmt.Println(" > Density:        ", solution.Density())
		fmt.Printf(" > Randomness:      %.1f%%\n", (100.0 * solution.Randomness()))
	}
	if puzzle != nil {
		undef := puzzle.CountSquares(model.SquareUndefined)
		all := puzzle.Size()
		fmt.Printf(" > Undefinedness:   %.1f%% (%v/%v)\n", (100.0 * float64(undef) / float64(all)), undef, all)

		difficulty := model.GetDifficulty(puzzle).String()
		fmt.Println(" > Difficulty:     ", difficulty)
	}
}
