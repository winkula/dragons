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
func Generate(width int, height int, difficulty Difficulty) *World {
	for {
		world := New(width, height).FillSquares(SquareEmpty)
		size := world.Width * world.Height
		numDragons := size / 6
		for i := 0; i < numDragons; {
			index := rand.Intn(size)
			if world.GetSquareByIndex(index) == SquareDragon {
				continue
			}
			worldNew := world.Clone()
			worldNew.SetDragon(index)
			if Validate(worldNew) {
				world = worldNew
			} else {
				continue
			}
			i++
		}
		if Validate(world) {
			return world
		}
	}
}

// GenerateFrom creates a puzzle from a given solved or partially solved puzzle and also takes a difficulty parameter.
func GenerateFrom(world *World, difficulty Difficulty) *World {
	bestWorld := world.Clone()
	mostUndefined := 0
	loops := 100
	tries := 100
	for i := 0; i < loops; i++ {
		w := incrementallyObfuscate(world, tries, difficulty)
		undefCount := w.CountSquares(SquareUndefined)
		if undefCount > mostUndefined {
			bestWorld = w
		}
	}
	return bestWorld
}

// incrementallyObfuscate takes a world state and incrementally sets squares to "undefined".
// After every step, it verifies if the puzzle is still solvable (i.e. has a distinct solution).
func incrementallyObfuscate(world *World, tries int, difficulty Difficulty) *World {
	if world.Size() == world.CountSquares(SquareUndefined) {
		panic("generateInternal: all squares undefined")
	}
	size := world.Size()
	for i := 0; i < tries; {
		index := rand.Intn(size)
		if world.GetSquareByIndex(index) == SquareUndefined {
			continue
		}
		worldNew := world.Clone()
		worldNew.SetSquareByIndex(index, SquareUndefined)
		if (difficulty > DifficultyEasy || len(EnumerateSquare(worldNew, index)) == 1) && HasDistinctSolution(worldNew) {
			world = worldNew
			i = 0
		}
		i++
	}
	return world
}
