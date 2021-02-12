package model

import (
	"testing"
)

func TestParseDifficulty(t *testing.T) {
	tables := []struct {
		input    string
		expected Difficulty
	}{
		{"", DifficultyUnknown},
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

		{"brutal", DifficultyBrutal},
		{"b", DifficultyBrutal},
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

func TestDifficultyString(t *testing.T) {
	tables := []struct {
		input    Difficulty
		expected string
	}{
		{DifficultyUnknown, "unknown"},
		{DifficultyEasy, "easy"},
		{DifficultyMedium, "medium"},
		{DifficultyHard, "hard"},
		{DifficultyBrutal, "brutal"},
	}

	for _, table := range tables {
		str := table.input.String()
		if str != table.expected {
			t.Errorf("String was incorrect, got: %v, want: %v.\n",
				str,
				table.expected)
		}
	}
}
