package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/winkula/dragons/pkg/model"
)

var generateCmd = flag.NewFlagSet("generate", flag.ExitOnError)
var genArgDifficulty = generateCmd.String("difficulty", "easy", "difficulty of the puzzle")
var genArgSize = generateCmd.Int("s", 8, "the puzzles size")
var genArgWidth = generateCmd.Int("w", 0, "the puzzles width")
var genArgHeight = generateCmd.Int("h", 0, "the puzzles height")
var genArgDuration = generateCmd.Duration("duration", 2*time.Second, "the available time to generate the puzzle")

func init() {
	registerCommand("generate", generateCmd, func() {
		difficultyEnum := model.ParseDifficulty(*genArgDifficulty)

		if *genArgWidth == 0 {
			*genArgWidth = *genArgSize
		}
		if *genArgHeight == 0 {
			*genArgHeight = *genArgSize
		}

		if len(generateCmd.Args()) > 0 {
			g := parse(generateCmd.Arg(0), false)
			generateFrom(g, difficultyEnum, *genArgDuration)
		} else {
			generate(*genArgWidth, *genArgHeight, difficultyEnum, *genArgDuration/2)
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
