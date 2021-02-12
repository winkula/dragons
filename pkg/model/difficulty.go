package model

import "strings"

// Difficulty represents possible difficulty levels.
type Difficulty int

const (
	// DifficultyUnknown represents unknown difficulty.
	DifficultyUnknown = iota
	// DifficultyEasy represents "easy" puzzles.
	DifficultyEasy
	// DifficultyMedium represents "medium" puzzles.
	DifficultyMedium
	// DifficultyHard represents "hard" puzzles.
	DifficultyHard
	// DifficultyBrutal represents "brutal" puzzles (only solvable using brute force).
	DifficultyBrutal
)

// ParseDifficulty takes a string and returns its Difficulty value.
func ParseDifficulty(str string) Difficulty {
	if len(str) == 0 {
		return DifficultyUnknown
	}

	str = strings.ToLower(str)
	if strings.HasPrefix("easy", str) {
		return DifficultyEasy
	}
	if strings.HasPrefix("medium", str) {
		return DifficultyMedium
	}
	if strings.HasPrefix("hard", str) {
		return DifficultyHard
	}
	if strings.HasPrefix("brutal", str) {
		return DifficultyBrutal
	}
	return DifficultyUnknown
}

func (difficulty Difficulty) String() string {
	return []string{"unknown", "easy", "medium", "hard", "brutal"}[difficulty]
}
