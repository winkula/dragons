package model

// Options defines possible square values when enumerating valid grids.
var options = []Square{
	SquareDragon,
	SquareFire,
	SquareEmpty,
}

type gridPredicate func(*Grid) bool
type gridsPredicate func([]*Grid) bool

// Enumerate enumerates all possible successors from a given grid.
func Enumerate(g *Grid) []*Grid {
	return enumerate(g, all, stopNever)
}

// EnumerateFilter enumerates all possible successors from a given grid and
// allows to filter only the wanted grids.
func EnumerateFilter(g *Grid, filter gridPredicate) []*Grid {
	return enumerate(g, filter, stopNever)
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
	return len(enumerate(g, all, stopWhenMultipleSolutions)) == 1
}

func enumerate(g *Grid, filter gridPredicate, isEarlyStop gridsPredicate) []*Grid {
	res := make([]*Grid, 0)
	if !Validate(g) {
		// skip further investigation because the state is invalid
		return res
	}
	return enumRecur(g, res, 0, filter, isEarlyStop)
}

func enumRecur(g *Grid, res []*Grid, i int, filter gridPredicate, isEarlyStop gridsPredicate) []*Grid {
	if isEarlyStop(res) {
		// early exit is used to validate, if a distinct solution exists
		return res
	}

	if isLeaf(g) {
		if filter(g) {
			// valid leaf node found, add grid to result
			res = append(res, g)
		}
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

			res = enumRecur(suc, res, i+1, filter, isEarlyStop)
		}
	} else {
		suc := g.Clone()
		res = enumRecur(suc, res, i+1, filter, isEarlyStop)
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

func all(g *Grid) bool {
	return true
}
