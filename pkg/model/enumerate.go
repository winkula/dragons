package model

import (
	"sort"
)

// Options defines possible square values when enumerating valid grids.
var options = []Square{
	SquareDragon,
	SquareFire,
	SquareEmpty,
}

// Enumerate enumerates all possible successors from a given grid.
func Enumerate(g *Grid) []*Grid {
	return enumerate(g, stopNever)
}

// EnumerateSquare enumerates all possible grids if a specified square is undefined.
// Otherwise, the given grid is returned as single possible solution.
func EnumerateSquare(g *Grid, i int) []*Grid {
	if g.Squarei(i) != SquareUndefined {
		return []*Grid{g.Clone()}
	}
	result := make([]*Grid, 0)
	for _, option := range options {
		suc := g.Clone()
		suc.SetSquarei(i, option)
		if Validate(suc) {
			result = append(result, suc)
		}
	}
	return result
}

// IsDistinct returns true if there is exactly one solution to the given grid.
func IsDistinct(g *Grid) bool {
	return len(enumerate(g, stopWhenMultipleSolutions)) == 1
}

// MostInteresting returns the most interesting possible solution to a puzzle.
// TODO: optimize this (this only works for small grids)
func MostInteresting(g *Grid) *Grid {
	sucs := Enumerate(g)
	byInterestingnessDesc := func(i, j int) bool {
		return sucs[i].Interestingness() > sucs[j].Interestingness()
	}
	// sort by interestingness
	sort.Slice(sucs, byInterestingnessDesc)
	return sucs[0]
}

func enumerate(g *Grid, isEarlyStop func([]*Grid) bool) []*Grid {
	res := make([]*Grid, 0)
	if !Validate(g) {
		// skip further investigation because the state is invalid
		return res
	}
	return enumRecur(g, res, 0, isEarlyStop)
}

func enumRecur(g *Grid, res []*Grid, i int, isEarlyStop func([]*Grid) bool) []*Grid {
	if isEarlyStop(res) {
		// early exit is used to validate, if a distinct solution exists
		return res
	}

	if isLeaf(g) {
		// valid leaf node found, add grid to result
		res = append(res, g)
		return res
	}

	square := g.Squarei(i)
	if square == SquareUndefined {
		for _, option := range options {
			suc := g.Clone()
			suc.SetSquarei(i, option)

			// check all neighbor square plus the square that was changed
			toCheck := append(suc.NeighborIndicesi(i, false), i)
			if !ValidatePartial(suc, toCheck) {
				continue
			}

			res = enumRecur(suc, res, i+1, isEarlyStop)
		}
	} else {
		suc := g.Clone()
		res = enumRecur(suc, res, i+1, isEarlyStop)
	}
	return res
}

func isLeaf(g *Grid) bool {
	return !g.HasSquare(SquareUndefined)
}

func stopWhenMultipleSolutions(res []*Grid) bool {
	return len(res) > 1
}

func stopNever(res []*Grid) bool {
	return false
}
