package model

type rule (func(*World) bool)

var rules = []rule{
	// Territory rule.
	// A dragon can not have other dragons in his territory.
	func(w *World) bool {
		for i, square := range w.Squares {
			if square == SquareDragon {
				x, y := w.GetCoords(i)
				neighbours := w.GetNeighbours(x, y)
				for _, n := range neighbours {
					if n == SquareDragon {
						return false
					}
				}
			}
		}
		return true
	},
	// Survive rule.
	// A dragon can not be surrounded by more than 50% fire (only counting fields on the grid).
	func(w *World) bool {
		for i, square := range w.Squares {
			if square == SquareDragon {
				total := 0
				fire := 0

				x, y := w.GetCoords(i)
				neighbours := w.GetNeighbours(x, y)
				for _, n := range neighbours {
					if n == SquareFire {
						fire++
					}
					if n != SquareOut {
						total++
					}
				}

				if 2*fire > total {
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
			x, y := w.GetCoords(i)
			neighbours := w.GetNeighbours(x, y)

			dragons := 0

			for _, n := range neighbours {
				if n == SquareDragon {
					dragons++
				}
			}

			if dragons > 1 && square != SquareFire {
				return false
			}
		}
		return true
	},
}

// ValidateWorld validates, if a world is in a valid state.
func ValidateWorld(w *World) bool {
	for _, rule := range rules {
		if !rule(w) {
			return false
		}
	}
	return true
}
