package model

type rule struct {
	name  string
	check func(w *World, i int) bool
}

// rules represents the rules that defines, if a dragons puzzle is valid or not.
var rules = []rule{
	// Dragons can not have other dragons in their territory.
	// Territory means the 8 squares that are around itself.
	{
		name: "territory",
		check: func(w *World, i int) bool {
			square := w.GetSquareByIndex(i)
			if square == SquareDragon {
				return w.CountNeighbours(i, SquareDragon) == 0
			}
			return true
		},
	},
	// Overlapping territories must be fire.
	// Every square that is part of multiple territories must be fire - and only then.
	{
		name: "fight",
		check: func(w *World, i int) bool {
			square := w.GetSquareByIndex(i)
			if square == SquareFire {
				// if a square is fire, there must be at least two dragons around it
				dragons := w.CountNeighbours(i, SquareDragon)
				undefined := w.CountNeighbours(i, SquareUndefined)
				return dragons+undefined > 1
			}
			if square == SquareEmpty {
				// if 0 or 1 dragon is around a square, there must NOT be fire
				// if a square is not fire, there can maximum be one dragon around it
				return w.CountNeighbours(i, SquareDragon) <= 1
			}
			return true
		},
	},
	// At least 2 of the adjacent squares of a dragon must be empty.
	// If a dragon is at the edge or in the corner of the grid, the possible number of adjacent squares is reduced.
	{
		name: "survive",
		check: func(w *World, i int) bool {
			square := w.GetSquareByIndex(i)
			if square == SquareDragon {
				empty := w.CountAdjacentNeighbours(i, SquareEmpty)
				undef := w.CountAdjacentNeighbours(i, SquareUndefined)
				return empty+undef >= 2
			}
			return true
		},
	},
}

// Validate validates, if a world is in a valid state.
// This applies all validation rules to the world state.
func Validate(w *World) bool {
	for i := range w.Squares {
		for _, rule := range rules {
			if !rule.check(w, i) {
				return false
			}
		}
	}
	return true
}
