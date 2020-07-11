package model

import (
	"math/rand"
)

// Difficulty represents possible difficulty levels.
type Difficulty int

const (
	// DifficultyEasy represents "easy" puzzles.
	DifficultyEasy = iota
	// DifficultyMedium represents "medium" puzzles.
	DifficultyMedium
	// DifficultyHard represents "hard" puzzles.
	DifficultyHard
)

// Generate generates a random puzzle with the given dimensions and difficulty.
// TODO: rework this implementation!
func Generate(width int, height int, difficulty Difficulty) *Grid {
	for {
		g := New(width, height).FillSquares(SquareEmpty)
		size := g.Width * g.Height
		numDragons := size / 6
		for i := 0; i < numDragons; {
			index := rand.Intn(size)
			if g.Squarei(index) == SquareDragon {
				continue
			}
			suc := g.Clone()
			suc.SetDragon(index)
			if Validate(suc) {
				g = suc
			} else {
				continue
			}
			i++
		}
		if Validate(g) {
			return g
		}
	}
}

// GenerateFrom creates a puzzle from a given solved or partially solved puzzle and also takes a difficulty parameter.
func GenerateFrom(g *Grid, difficulty Difficulty) *Grid {
	best := g.Clone()
	mostUndefined := 0
	loops := 100
	tries := 100
	for i := 0; i < loops; i++ {
		suc := incrementallyObfuscate(g, tries, difficulty)
		undefCount := suc.CountSquares(SquareUndefined)
		if undefCount > mostUndefined {
			best = suc
		}
	}
	return best
}

// incrementallyObfuscate takes a grid state and incrementally sets squares to "undefined".
// After every step, it verifies if the puzzle is still solvable (i.e. has a distinct solution).
func incrementallyObfuscate(g *Grid, tries int, difficulty Difficulty) *Grid {
	if g.Size() == g.CountSquares(SquareUndefined) {
		panic("generateInternal: all squares undefined")
	}
	size := g.Size()
	for i := 0; i < tries; {
		index := rand.Intn(size)
		if g.Squarei(index) == SquareUndefined {
			continue
		}
		suc := g.Clone()
		suc.SetSquarei(index, SquareUndefined)
		if (difficulty > DifficultyEasy || len(EnumerateSquare(suc, index)) == 1) && HasDistinctSolution(suc) {
			g = suc
			i = 0
		}
		i++
	}
	return g
}
