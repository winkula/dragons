package model

import (
	"sort"
)

// Options defines possible square values when enumerating valid grids.
var options = []Square{SquareDragon, SquareFire, SquareEmpty}

// Enumerate enumerates all possible successors from a given grid.
func Enumerate(g *Grid) []*Grid {
	return enumerateInternal(g, make([]*Grid, 0), 0, false)
}

// EnumerateSquare enumerates all possible grids if a specified square is undefined.
// Otherwise, the given grid is returned as single possible solution.
func EnumerateSquare(g *Grid, index int) []*Grid {
	if g.Squarei(index) != SquareUndefined {
		return []*Grid{g.Clone()}
	}
	result := make([]*Grid, 0)
	for _, option := range options {
		suc := g.Clone()
		suc.SetSquarei(index, option)
		if Validate(suc) {
			result = append(result, suc)
		}
	}
	return result
}

// HasDistinctSolution returns 'true' if there is exactly one solution to the problem.
func HasDistinctSolution(g *Grid) bool {
	return len(enumerateInternal(g, make([]*Grid, 0), 0, true)) == 1
}

// MostInteresting returns the most interesting possible solution to a puzzle.
// TODO: optimize this (this only works for small grids)
func MostInteresting(g *Grid) *Grid {
	sucs := Enumerate(g)
	// sort by interestingness
	byInterestingnessDesc := func(i, j int) bool {
		return sucs[i].Interestingness() > sucs[j].Interestingness()
	}
	sort.Slice(sucs, byInterestingnessDesc)
	return sucs[0]
}

func hasSquares(g *Grid, square Square) bool {
	for _, v := range g.Squares {
		if v == square {
			return true
		}
	}
	return false
}

func enumerateInternal(g *Grid, result []*Grid, index int, skipOnMultipleSolutions bool) []*Grid {
	if skipOnMultipleSolutions && len(result) > 1 {
		// this early exit is used to validate, if a distinct solution exists
		return result
	}

	if !Validate(g) {
		// skip further investigation because the state is invalid
		// logd.Printf("%vInvalid!\n", strings.Repeat(" ", index))
		// debug(w)
		return result
	}

	if !hasSquares(g, SquareUndefined) {
		// only append, if its a leaf node (final state, no undefined squares)
		// debug("Sucessor found!")
		result = append(result, g)
		return result
	}

	if index >= len(g.Squares) {
		panic("Force recursion stop (this should never happen)!!!")
	}

	square := g.Squarei(index)
	if square == SquareUndefined {
		for _, option := range options {
			//logd.Printf("%vOption %c for x=%v,y=%v\n", strings.Repeat(" ", index), squareSymbols[option], x, y)
			suc := g.Clone()
			suc.SetSquarei(index, option)
			result = enumerateInternal(suc, result, index+1, skipOnMultipleSolutions)
		}
	} else {
		//logd.Printf("%vSquare already defined for i=%v x=%v,y=%v\n", strings.Repeat(" ", index), index, x, y)
		suc := g.Clone()
		result = enumerateInternal(suc, result, index+1, skipOnMultipleSolutions)
	}
	return result
}
