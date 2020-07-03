package model

// Enumerate enumerates all possible successor states from a given world.
func (w *World) Enumerate() []*World {
	return enumerateInternal(w, make([]*World, 0), 0)
}

func hasSquaresOfType(w *World, square Square) bool {
	for _, v := range w.Squares {
		if v == square {
			return true
		}
	}
	return false
}

func enumerateInternal(w *World, result []*World, index int) []*World {
	if index >= len(w.Squares) {
		return result // stop recursion
	}

	x, y := w.GetCoords(index)
	square := w.GetSquare(x, y)
	if square != SquareUndefined {
		return result // square already defined
	}

	//fmt.Printf("Enumerate i=%v, x=%v, y=%v", i, x, y)
	for _, option := range Options {
		// fmt.Printf("Enumerate (option %v) for world: %s", option, w)
		successor := w.Clone()
		successor.SetSquare(x, y, option)
		if ValidateWorld(successor) {
			if !hasSquaresOfType(successor, SquareUndefined) {
				// only append, if its a leaf node (final state, no undefined squares)
				result = append(result, successor)
			}
			result = enumerateInternal(successor, result, index+1)
		}
	}
	return result
}
