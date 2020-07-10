package model

import "testing"

func TestValidate(t *testing.T) {
	tables := []struct {
		world *World
		valid bool
		rule  string
	}{
		// The territory rule is fulfilled.
		{New(3, 3), true, "territory"},
		{New(3, 3).
			SetSquare(0, 0, SquareDragon), true, "territory"},
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(1, 1, SquareFire).
			SetSquare(2, 2, SquareDragon), true, "territory"},
		// The territory rule is violated.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(1, 1, SquareDragon), false, "territory"},
		{New(3, 3).
			SetSquare(1, 0, SquareDragon).
			SetSquare(2, 0, SquareDragon), false, "territory"},
		// The survive rule is fulfilled.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon), true, "survive"},
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(1, 1, SquareFire), true, "survive"},
		// The survive rule is violated.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(1, 0, SquareFire).
			SetSquare(0, 1, SquareFire), false, "survive"},
		{New(3, 3).
			SetSquare(1, 1, SquareDragon).
			SetSquare(0, 0, SquareFire).
			SetSquare(1, 0, SquareFire).
			SetSquare(2, 0, SquareFire).
			SetSquare(0, 1, SquareFire).
			SetSquare(2, 1, SquareFire), false, "survive"},
		// The fight rule is fulfilled.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(2, 2, SquareDragon).
			SetSquare(1, 1, SquareFire), true, "fight"},
		// The fight rule is violated.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(2, 1, SquareDragon).
			SetSquare(1, 0, SquareEmpty).
			SetSquare(1, 1, SquareFire), false, "fight"},
		{New(1, 1).
			SetSquare(0, 0, SquareFire), false, "fight"},
	}

	for _, table := range tables {
		valid := Validate(table.world)
		if valid != table.valid {
			t.Errorf("ValidateWorld was incorrect, got: %t, want: %t (%s rule). World: \n%s",
				valid,
				table.valid,
				table.rule,
				table.world)
		}
	}
}
