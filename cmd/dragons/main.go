package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/winkula/dragons/pkg/model"
)

func main() {
	seed()

	difficulty := flag.String("difficulty", "easy", "difficulty of the puzzle")
	flag.Parse()

	tail := flag.Args()
	cmd := tail[0]
	difficultyEnum := parseDifficulty(*difficulty)

	if cmd == "parse" {
		parse(tail[1], true)
		return
	}

	if cmd == "enum" {
		g := parse(tail[1], true)
		enumerate(g)
		return
	}

	if cmd == "gen" {
		g := parse(tail[1], false)
		generatePuzzle(g, difficultyEnum)
		return
	}

	if cmd == "gens" {
		g := parse(tail[1], false)
		generateSolution(g)
		return
	}

	if cmd == "genr" {
		g := parse(tail[1], false)
		generateRandom(g.Width, g.Height, difficultyEnum)
		return
	}

	if cmd == "solve" {
		g := parse(tail[1], true)
		solve(g, difficultyEnum)
		return
	}

	if cmd == "val" {
		g := parse(tail[1], true)
		validate(g)
		return
	}
}

func parse(s string, print bool) *model.Grid {
	g := model.Parse(s)
	if print {
		fmt.Println("Grid:")
		fmt.Println(g)
	}
	return g
}

func enumerate(g *model.Grid) {
	sucs := model.Enumerate(g)
	sort.Slice(sucs, func(i, j int) bool {
		return sucs[i].Interestingness() < sucs[j].Interestingness()
	})
	for _, s := range sucs {
		fmt.Println(s)
		fmt.Println("Interestingness:", s.Interestingness())
	}
	fmt.Printf("%v possible grids found.\n", len(sucs))
}

func generateSolution(g *model.Grid) {
	if g.Size() == g.CountSquares(model.SquareUndefined) {
		g = model.MostInteresting(g)
	}
	fmt.Println("Input:")
	fmt.Println(g)

	solution := model.GenerateFrom(g, model.DifficultyEasy)
	fmt.Println("Solution:")
	fmt.Println(solution)

	printStats(nil, solution)
}

func generatePuzzle(g *model.Grid, difficulty model.Difficulty) {
	fmt.Println("Input:")
	fmt.Println(g)

	solution := model.Generate(g.Width, g.Height)
	fmt.Println("Solution:")
	fmt.Println(solution)

	puzzle := model.GenerateFrom(solution, difficulty)
	fmt.Println("Puzzle:")
	fmt.Println(puzzle)

	printStats(puzzle, solution)
}

func generateRandom(width int, height int, difficulty model.Difficulty) {
	solution := model.Generate(width, height)
	fmt.Println("Solution:")
	fmt.Println(solution)

	puzzle := model.GenerateFrom(solution, difficulty)
	fmt.Println("Puzzle:")
	fmt.Println(puzzle)

	printStats(puzzle, solution)
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

func validate(g *model.Grid) {
	valid := model.Validate(g)
	fmt.Println("Valid:", valid)
}

func seed() {
	// seed the random generator
	rand.Seed(time.Now().UnixNano())
}

func parseDifficulty(str string) model.Difficulty {
	switch str {
	case "easy":
		return model.DifficultyEasy
	case "medium":
		return model.DifficultyMedium
	case "hard":
		return model.DifficultyHard
	default:
		return model.DifficultyEasy
	}
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
