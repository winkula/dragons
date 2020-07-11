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

// Generate generates a random puzzle with the given dimensions and difficulty.
// TODO: rework this implementation!
func Generate(width int, height int, difficulty Difficulty) *Grid {
	duration := 5.0 // the generation can last this many seconds
	maxFails := 10000
	start := time.Now()
	var best *Grid
	for {
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
		if time.Since(start).Seconds() > duration {
			break
		}
	}
	return best
}

// GenerateFrom creates a puzzle from a given solved or partially solved puzzle and also takes a difficulty parameter.
func GenerateFrom(g *Grid, difficulty Difficulty) *Grid {
	best := g.Clone()
	mostUndefined := 0
	loops := 100
	tries := 100
	for i := 0; i < loops; i++ {
		suc := obfuscIncr(g, tries, difficulty)
		undefCount := suc.CountSquares(SquareUndefined)
		if undefCount > mostUndefined {
			best = suc
		}
	}
	return best
}

// obfuscIncr takes a grid state and incrementally sets squares to "undefined".
// After every step, it verifies if the puzzle is still solvable (i.e. has a distinct solution).
func obfuscIncr(g *Grid, tries int, difficulty Difficulty) *Grid {
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
		if (difficulty > DifficultyEasy || len(EnumerateSquare(suc, index)) == 1) && IsDistinct(suc) {
			g = suc
			i = 0
		}
		i++
	}
	return g
}
