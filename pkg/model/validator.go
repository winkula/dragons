package model

type rule struct {
	name    string
	squares Square
	check   func(w *World, i int) bool
	apply   func(w *World, i int) *World
}

var rules = []rule{
	// A dragon can not have other dragons in his territory.
	rule{
		name:    "territory",
		squares: SquareDragon,
		check: func(w *World, i int) bool {
			square := w.GetSquareByIndex(i)
			if square == SquareDragon {
				return w.CountNeighbours(i, SquareDragon) == 0
			}
			return true
		},
		apply: func(w *World, i int) *World {
			square := w.GetSquareByIndex(i)
			if square == SquareDragon {
				// TODO: mark all neightbour squares as "no dragon"
			}
			return w
		},
	},
	// Overlapping territories must be fire.
	rule{
		name:    "fight",
		squares: SquareFire | SquareEmpty,
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
		apply: func(w *World, i int) *World {
			// TODO
			return w
		},
	},
	// At least 2 of the adjacent squares of a dragon must be empty.
	rule{
		name:    "survive",
		squares: SquareDragon,
		check: func(w *World, i int) bool {
			square := w.GetSquareByIndex(i)
			if square == SquareDragon {
				empty := w.CountAdjacentNeighbours(i, SquareEmpty)
				undef := w.CountAdjacentNeighbours(i, SquareUndefined)
				return empty+undef >= 2
			}
			return true
		},
		apply: func(w *World, i int) *World {
			// TODO
			return w
		},
	},
}

/*
type rule (func(*World) bool)

var rules = []rule{
	// Territory rule.
	// A dragon can not have other dragons in his territory.
	func(w *World) bool {
		for i, square := range w.Squares {
			if square == SquareDragon {
				dragons := w.CountNeighbours(i, SquareDragon)
				if dragons > 0 {
					return false
				}
			}
		}
		return true
	},
	// Fight rule.
	// Overlapping territories must be fire.
	func(w *World) bool {
		for i, square := range w.Squares {
			if square != SquareUndefined {
				dragons := w.CountNeighbours(i, SquareDragon)
				undefined := w.CountNeighbours(i, SquareUndefined)
				if square != SquareFire && square != SquareUndefined && dragons > 1 {
					// if more than one dragon is around a square, there must be fire
					return false
				}
				if square == SquareFire && dragons <= 1 && undefined == 0 {
					// if 0 or 1 dragon is around a square, there must NOT be fire
					return false
				}
			}
		}
		return true
	},
	// Escape rule.
	func(w *World) bool {
		for i, square := range w.Squares {
			if square == SquareDragon {
				empty := w.CountAdjacentNeighbours(i, SquareEmpty)
				undef := w.CountAdjacentNeighbours(i, SquareUndefined)

				if empty+undef < 2 {
					return false
				}
			}
		}
		return true
	},
	// Original survive rule.
	// A dragon can not be surrounded by more than 50% fire (only counting fields on the grid).
	func(w *World) bool {
		for i, square := range w.Squares {
			if square == SquareDragon {
				total := 8 - w.CountNeighbours(i, SquareOut)
				fire := w.CountNeighbours(i, SquareFire)

				if 2*fire > total {
					return false
				}
			}
		}
		return true
	},
	// The "fire around empty squares" rule.
	func(w *World) bool {
		for i, square := range w.Squares {
			if square == SquareEmpty {
				fire := w.CountNeighbours(i, SquareFire)
				undefined := w.CountNeighbours(i, SquareUndefined)

				fireMin := 1
				fireMax := math.MaxInt32

				if fire+undefined < fireMin {
					return false // too few fires
				}
				if fire > fireMax {
					return false // to much fire
				}
			}
		}
		return true
	},
}
*/

// ValidateWorld validates, if a world is in a valid state.
func ValidateWorld(w *World) bool {
	for i := range w.Squares {
		for _, rule := range rules {
			if !rule.check(w, i) {
				return false
			}
		}
	}
	return true
}
