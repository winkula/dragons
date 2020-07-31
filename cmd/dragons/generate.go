package main

import (
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

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
