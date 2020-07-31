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

func generate(width int, height int, difficulty model.Difficulty) {
	solution := model.Generate(width, height)
	generateFrom(solution, difficulty)
}

func generateFrom(solution *model.Grid, difficulty model.Difficulty) {
	if solution.IsUndefined() {
		generate(solution.Width, solution.Height, difficulty)
		return
	}

	fmt.Println("Solution:")
	fmt.Println(solution)

	puzzle := model.GenerateFrom(solution, difficulty, 5)
	fmt.Println("Puzzle:")
	fmt.Println(puzzle)

	printStats(puzzle, solution)
}
