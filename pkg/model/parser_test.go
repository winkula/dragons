package model

import (
	"reflect"
	"testing"
)

func TestParseWorld(t *testing.T) {
	tables := []struct {
		str   string
		world *World
	}{
		{"_", NewWorld(1, 1)},
		{"___,___", NewWorld(3, 2)},
		{"_D", NewWorld(2, 1).SetSquare(1, 0, SquareDragon)},
	}

	for _, table := range tables {
		world := ParseWorld(table.str)
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
	ParseWorld("__,___")
}

func TestParseWorldPanicsOnInvalidSymbols(t *testing.T) {
	defer assertPanic(t)
	ParseWorld("ab,xy,+*")
}

func assertPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic.")
	}
}
