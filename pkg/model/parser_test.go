package model

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tables := []struct {
		str   string
		world *World
	}{
		{"_", New(1, 1)},
		{"___,___", New(3, 2)},
		{"_D", New(2, 1).SetSquare(1, 0, SquareDragon)},

		{"1,1", New(1, 1)},
		{"2,1", New(2, 1)},
		{"4,5", New(4, 5)},
	}

	for _, table := range tables {
		world := Parse(table.str)
		if !reflect.DeepEqual(world, table.world) {
			t.Errorf("ParseWorld was incorrect, got:\n%s, want:\n%s.\nInput: '%s'",
				world,
				table.world,
				table.str)
		}
	}
}

func TestParseWorldPanicsOnWrongDimensions(t *testing.T) {
	defer assertPanic(t)
	Parse("__,___")
}

func TestParseWorldPanicsOnInvalidSymbols(t *testing.T) {
	defer assertPanic(t)
	Parse("ab,xy,+*")
}

func assertPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic.")
	}
}
