package model

import (
	"testing"
)

func TestParseDifficulty(t *testing.T) {
	tables := []struct {
		input    string
		expected Difficulty
	}{
		{"  ", DifficultyUnknown},
		{"?*%*ç", DifficultyUnknown},

		{"easy", DifficultyEasy},
		{"EASY", DifficultyEasy},
		{"e", DifficultyEasy},

		{"medium", DifficultyMedium},
		{"med", DifficultyMedium},
		{"m", DifficultyMedium},

		{"hard", DifficultyHard},
		{"h", DifficultyHard},
	}

	for _, table := range tables {
		d := ParseDifficulty(table.input)
		if d != table.expected {
			t.Errorf("ParseDifficulty was incorrect, got: %v, want: %v.\n",
				d,
				table.expected)
		}
	}
}
