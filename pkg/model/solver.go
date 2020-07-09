package model

// Solve solves a world and returns the only possible solution or otherwise panics.
func Solve(w *World) *World {
	knowledge := newKnowledge(w)
	for i := 0; i < 100; i++ {
		for i := range w.Squares {
			for _, r := range solveRules {
				r(w, i, knowledge)
				if isSolved(w) {
					return w
				}
			}
		}
	}
	panic("Solver: No solution found!")
}

func isSolved(w *World) bool {
	return w.CountSquares(SquareUndefined) == 0 && ValidateWorld(w)
}

type knowledge []uint8

func newKnowledge(w *World) knowledge {
	k := make(knowledge, w.Size())
	all := uint8(1<<SquareDragon | 1<<SquareFire | 1<<SquareEmpty)
	for i := range w.Squares {
		k[i] = all
	}
	return k
}

func (k knowledge) squareCannotBe(i int, v Square) {
	k[i] &= ^(1 << i)
}

type solveRule (func(*World, int, knowledge) *World)

var solveRules = []solveRule{
	// if a dragon is set, no neighbour square can be a dragon too
	func(w *World, i int, k knowledge) *World {
		square := w.GetSquareByIndex(i)
		if square == SquareDragon {
			for _, ni := range w.GetNeighbourIndexes(i) {
				k.squareCannotBe(ni, SquareDragon)
			}
		}
		return w
	},

	// if a dragon is set, at least two adjacent quares must be empty
	func(w *World, i int, k knowledge) *World {
		square := w.GetSquareByIndex(i)
		if square == SquareDragon {
			empty := w.CountAdjacentNeighbours(i, SquareEmpty)
			undef := w.CountAdjacentNeighbours(i, SquareEmpty)
			if empty < 2 && empty+undef == 2 {
				for _, ni := range w.GetAdjacentNeighbourIndexes(i) {
					if w.GetSquareByIndex(ni) == SquareUndefined {
						w.SetSquareByIndex(ni, SquareEmpty)
					}
				}
			}
		}
		return w
	},

	// if a fire square is set, there must be at least 2 dragons around it
	func(w *World, i int, k knowledge) *World {
		square := w.GetSquareByIndex(i)
		if square == SquareDragon {
			dragons := w.CountNeighbours(i, SquareDragon)
			undef := w.CountNeighbours(i, SquareEmpty)
			if dragons < 2 && dragons+undef == 2 {
				for _, ni := range w.GetNeighbourIndexes(i) {
					if w.GetSquareByIndex(ni) == SquareUndefined {
						w.SetSquareByIndex(ni, SquareDragon)
					}
				}
			}
		}
		return w
	},

	// if a undefined square is surrounded by more than one dragon, there must be fire
	func(w *World, i int, k knowledge) *World {
		square := w.GetSquareByIndex(i)
		if square == SquareUndefined {
			dragons := w.CountNeighbours(i, SquareDragon)
			if dragons > 1 {
				w.SetSquareByIndex(i, SquareFire)
			}
		}
		return w
	},

	// if a empty square is set, there can maximum be one dragon around it
	func(w *World, i int, k knowledge) *World {
		square := w.GetSquareByIndex(i)
		if square == SquareEmpty {
			dragons := w.CountNeighbours(i, SquareDragon)
			if dragons == 1 {
				for _, ni := range w.GetNeighbourIndexes(i) {
					k.squareCannotBe(ni, SquareDragon)
				}
			}
		}
		return w
	},
}
