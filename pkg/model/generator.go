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

// GenerateWorldOld generates a random world with the given dimensions.
func GenerateWorldOld(width int, height int) *World {
	for {
		world := NewWorld(width, height).FillSquares(SquareEmpty)
		size := world.Width * world.Height
		numDragons := size / 6
		for i := 0; i < numDragons; {
			index := rand.Intn(size)
			if world.GetSquareByIndex(index) == SquareDragon {
				continue
			}
			worldNew := world.Clone()
			worldNew.SetDragon(index)
			if ValidateWorld(worldNew) {
				world = worldNew
			} else {
				continue
			}
			i++
		}
		if ValidateWorld(world) {
			return world
		}
	}
}

// ObfuscateWorld takes a world and tries to obfuscate it so that some
// information is missing but it still has a definite solution.
func ObfuscateWorld(world *World) *World {
	size := world.Size()
	tries := 20
	for i := 0; i < tries; {
		index := rand.Intn(size)
		if world.GetSquareByIndex(index) == SquareUndefined {
			continue
		}
		worldNew := world.Clone()
		worldNew.SetSquareByIndex(index, SquareUndefined)
		if worldNew.HasDistinctSolution() {
			world = worldNew
			i = 0
		}
		i++
	}
	return world
}

// GenerateWorld creates a puzzle with a distinct solution from an existing world.
func GenerateWorld(world *World, difficulty Difficulty) *World {
	bestWorld := world.Clone()
	mostUndefined := 0
	loops := 100
	tries := 100
	for i := 0; i < loops; i++ {
		w := generateInternal(world, tries, difficulty)
		undefCount := w.CountSquares(SquareUndefined)
		if undefCount > mostUndefined {
			bestWorld = w
		}
	}
	return bestWorld
}

func generateInternal(world *World, tries int, difficulty Difficulty) *World {
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
		if (difficulty > DifficultyEasy || len(worldNew.EnumerateSquare(index)) == 1) && worldNew.HasDistinctSolution() {
			world = worldNew
			i = 0
		}
		i++
	}
	return world
}
