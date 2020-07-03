package model

import "testing"

func TestEnumerate(t *testing.T) {
	tables := []struct {
		world *World
		count int
	}{
		{NewWorld(1, 1), 2},
		{NewWorld(1, 2), 3},
		{NewWorld(2, 1), 3},
		{NewWorld(2, 2), 5},
		{NewWorld(2, 3), 7},
		{NewWorld(3, 2), 7},
		{NewWorld(3, 3), 12},
	}

	for _, table := range tables {
		successors := table.world.Enumerate()
		length := len(successors)
		if length != table.count {
			t.Errorf("Enumerate was incorrect, got: %d, want: %d. World: \n%s",
				length,
				table.count,
				table.world)
		}
	}
}
