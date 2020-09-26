package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/winkula/dragons/pkg/model"
)

func init() {
	cmd := flag.NewFlagSet("generate", flag.ExitOnError)
	difficulty := cmd.String("difficulty", "easy", "difficulty of the puzzle")
	size := cmd.Int("s", 8, "the puzzles size")
	width := cmd.Int("w", 0, "the puzzles width")
	height := cmd.Int("h", 0, "the puzzles height")
	duration := cmd.Duration("duration", 2*time.Second, "the available time to generate the puzzle")

	registerCommand("generate", cmd, func() {
		difficultyEnum := model.ParseDifficulty(*difficulty)

		if *width == 0 {
			*width = *size
		}
		if *height == 0 {
			*height = *size
		}

		if len(cmd.Args()) > 0 {
			g := parse(cmd.Arg(0), false)
			generateFrom(g, difficultyEnum, *duration)
		} else {
			generate(*width, *height, difficultyEnum, *duration/2)
		}
	})
}

func generate(width int, height int, difficulty model.Difficulty, duration time.Duration) {
	solution := model.Generate(width, height, duration)
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

	printStats(puzzle, solution)
}
