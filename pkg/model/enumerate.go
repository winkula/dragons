package model

import "sort"

// Options defines possible square values when enumerating valid world states.
var options = []Square{SquareDragon, SquareFire, SquareEmpty}

// Enumerate enumerates all possible successor states from a given world.
func (w *World) Enumerate() []*World {
	return enumerateInternal(w, make([]*World, 0), 0, false)
}

// EnumerateSquare enumerates all possible worlds if a specified square is undefined.
// Otherwise, the given world is returned as single possible solution.
func (w *World) EnumerateSquare(index int) []*World {
	if w.GetSquareByIndex(index) != SquareUndefined {
		return []*World{w.Clone()}
	}
	result := make([]*World, 0)
	for _, option := range options {
		successor := w.Clone()
		successor.SetSquareByIndex(index, option)
		if ValidateWorld(successor) {
			result = append(result, successor)
		}
	}
	return result
}

// HasDistinctSolution returns 'true' if there is exactly one solution to the problem.
func (w *World) HasDistinctSolution() bool {
	return len(enumerateInternal(w, make([]*World, 0), 0, true)) == 1
}

// MostInteresting returns the most interesting possible solution to a puzzle.
func MostInteresting(puzzle *World) *World {
	successors := puzzle.Enumerate()
	// sort by interestingness
	byInterestingnessDesc := func(i, j int) bool {
		return successors[i].Interestingness() > successors[j].Interestingness()
	}
	sort.Slice(successors, byInterestingnessDesc)
	return successors[0]
}

func hasSquares(w *World, square Square) bool {
	for _, v := range w.Squares {
		if v == square {
			return true
		}
	}
	return false
}

func enumerateInternal(w *World, result []*World, index int, skipOnMultipleSolutions bool) []*World {
	if skipOnMultipleSolutions && len(result) > 1 {
		// this early exit is used to validate, if a distinct solution exists
		return result
	}

	if !ValidateWorld(w) {
		// skip further investigation because the state is invalid
		//fmt.Printf("%vInvalid!\n", strings.Repeat(" ", index))
		//fmt.Println(w)
		return result
	}

	if !hasSquares(w, SquareUndefined) {
		// only append, if its a leaf node (final state, no undefined squares)
		//fmt.Println("Sucessor found!")
		result = append(result, w)
		return result
	}

	if index >= len(w.Squares) {
		panic("Force recursion stop (this should never happen)!!!")
	}

	x, y := w.GetCoords(index)
	square := w.GetSquare(x, y)
	if square == SquareUndefined {
		for _, option := range options {
			//fmt.Printf("%vOption %c for x=%v,y=%v\n", strings.Repeat(" ", index), squareSymbols[option], x, y)
			successor := w.Clone()
			successor.SetSquare(x, y, option)
			result = enumerateInternal(successor, result, index+1, skipOnMultipleSolutions)
		}
	} else {
		//fmt.Printf("%vSquare already defined for i=%v x=%v,y=%v\n", strings.Repeat(" ", index), index, x, y)
		successor := w.Clone()
		result = enumerateInternal(successor, result, index+1, skipOnMultipleSolutions)
	}
	return result
}
