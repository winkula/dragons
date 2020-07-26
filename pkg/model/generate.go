package model

import (
	"math/rand"
	"time"
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

// Generate generates a random but solved grid with the given dimensions.
// TODO: rework this implementation!
func Generate(width int, height int) *Grid {
	duration := 5.0 // the generation can last this many seconds
	maxFails := 10000
	var best *Grid

	executeTimeBound(duration, func() {
		g := New(width, height).Fill(SquareEmpty)
		size := g.Size()
		fails := 0
		for fails < maxFails {
			i := rand.Intn(size)
			if g.Squarei(i) != SquareEmpty || g.CountNeighbors(i, SquareDragon) > 0 {
				fails++
				continue
			}
			suc := g.Clone().SetDragon(i)
			if !Validate(suc) {
				fails++
				continue
			}
			g = suc
		}
		if best == nil || g.Interestingness() > best.Interestingness() {
			best = g
		}
	})

	return best
}

// GenerateFrom creates a puzzle from a given solved or partially solved puzzle and also takes a difficulty parameter.
func GenerateFrom(g *Grid, difficulty Difficulty) *Grid {
	if g.Size() == g.CountSquares(SquareUndefined) {
		panic("generateInternal: all squares undefined")
	}

	duration := 5.0 // the generation can last this many seconds
	best := g.Clone()
	mostUndefined := 0
	tries := 100

	executeTimeBound(duration, func() {
		suc := obfuscate(g, tries, difficulty)
		undefCount := suc.CountSquares(SquareUndefined)
		if undefCount > mostUndefined {
			best = suc
		}
	})

	return best
}

// obfuscate takes a grid state and incrementally sets squares to "undefined".
// This is done choosing random squares (with a timeout of n tries).
// After every step, it verifies if the puzzle is still solvable (i.e. has a distinct solution).
func obfuscate(g *Grid, tries int, difficulty Difficulty) *Grid {
	size := g.Size()
	for i := 0; i < tries; {
		index := rand.Intn(size)
		if g.Squarei(index) == SquareUndefined {
			continue
		}
		suc := g.Clone()
		suc.SetSquarei(index, SquareUndefined)
		if checkSolvable(suc, index, difficulty) {
			g = suc
			i = 0
		}
		i++
	}
	return g
}

func checkSolvable(g *Grid, index int, difficulty Difficulty) bool {
	if !IsDistinct(g) {
		return false
	}

	if difficulty == DifficultyHard {
		return true
	}

	solved, _ := SolveHuman(g, difficulty)
	if solved == nil {
		return false
	}

	return true
}

func executeTimeBound(timeout float64, action func()) {
	start := time.Now()
	for {
		action()
		if time.Since(start).Seconds() > timeout {
			break
		}
	}
}
