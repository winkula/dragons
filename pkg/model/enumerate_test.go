package model

import "testing"

func TestEnumerate(t *testing.T) {
	tables := []struct {
		world *World
		count int
	}{
		{NewWorld(1, 1), 1},
		{NewWorld(1, 2), 1},
		{NewWorld(2, 1), 1},
		{NewWorld(2, 2), 5},
		{NewWorld(2, 3), 7},
		{NewWorld(3, 2), 7},
		{NewWorld(3, 3), 14},

		{ParseWorld("____,_ff_,_xd_"), 2},
		{ParseWorld("____,_ffx,_xd_"), 1},
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

func TestEnumerateSquare(t *testing.T) {
	tables := []struct {
		world *World
		index int
		count int
	}{
		{ParseWorld("x_,xd"), 1, 1},
		{ParseWorld("xx,x_"), 3, 2},
	}

	for _, table := range tables {
		successors := table.world.EnumerateSquare(table.index)
		length := len(successors)
		if length != table.count {
			t.Errorf("Enumerate was incorrect, got: %d, want: %d. World: \n%s",
				length,
				table.count,
				table.world)
		}
	}
}
