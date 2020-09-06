package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

var generateCmd = flag.NewFlagSet("generate", flag.ExitOnError)
var genArgDifficulty = generateCmd.String("difficulty", "easy", "difficulty of the puzzle")
var genArgWidth = generateCmd.Int("w", 3, "the puzzles width")
var genArgHeight = generateCmd.Int("h", 3, "the puzzles height")
var genArgDuration = generateCmd.Float64("duration", 10, "the available time to generate the puzzle")

func generate(width int, height int, difficulty model.Difficulty, duration float64) {
	solution := model.Generate(width, height, duration)
	generateFrom(solution, difficulty, duration)
}

func generateFrom(solution *model.Grid, difficulty model.Difficulty, duration float64) {
	if solution.IsUndefined() {
		generate(solution.Width, solution.Height, difficulty, duration)
		return
	}

	fmt.Println("Solution:")
	fmt.Println(solution)

	puzzle := model.GenerateFrom(solution, difficulty, duration)
	fmt.Println("Puzzle:")
	fmt.Println(puzzle)

	printStats(puzzle, solution)
}
