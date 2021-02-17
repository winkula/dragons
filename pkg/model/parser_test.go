package model

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tables := []struct {
		str  string
		grid *Grid
	}{
		{"_", New(1, 1)},
		{"___,___", New(3, 2)},
		{"_D", New(2, 1).SetSquare(1, 0, SquareDragon)},
		{"_n", New(2, 1).SetSquare(1, 0, SquareNoDragon)},

		{"1,1", New(1, 1)},
		{"2,1", New(2, 1)},
		{"4,5", New(4, 5)},
	}

	for _, table := range tables {
		g := Parse(table.str)
		if !reflect.DeepEqual(g, table.grid) {
			t.Errorf("Parse was incorrect, got:\n%s, want:\n%s.\nInput: '%s'",
				g,
				table.grid,
				table.str)
		}
	}
}

func TestParsePanicsOnWrongDimensions(t *testing.T) {
	defer assertPanic(t)
	Parse("__,___")
}

func TestParsePanicsOnInvalidSymbols(t *testing.T) {
	defer assertPanic(t)
	Parse("ab,xy,+*")
}

func assertPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic.")
	}
}
