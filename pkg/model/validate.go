package model

type rule struct {
	name  string
	check func(g *Grid, i int) bool
	apply func(g *Grid, i int, s Square)
}

// rules represents the rules that defines, if a dragons puzzle is valid or not.
var rules = []rule{
	// Dragons can not have other dragons in their territory.
	// Territory means the 8 squares that are around itself.
	{
		name: "territory",
		check: func(g *Grid, i int) bool {
			square := g.Squarei(i)
			if square == SquareDragon {
				dragons := g.NeighborCount(i, SquareDragon, false, false)
				return dragons == 0
			}
			return true
		},
	},
	// Overlapping territories must be fire.
	// Every square that is part of multiple territories must be fire - and only then.
	{
		name: "fight",
		check: func(g *Grid, i int) bool {
			square := g.Squarei(i)
			if square == SquareFire {
				// if a square is fire, there must be at least two dragons around it
				possibleDragons := g.NeighborCount(i, SquareDragon, false, true)
				return possibleDragons > 1
			}
			if square == SquareEmpty {
				// if 0 or 1 dragon is around a square, there must NOT be fire
				// if a square is not fire, there can maximum be one dragon around it
				dragons := g.NeighborCount(i, SquareDragon, false, false)
				return dragons <= 1
			}
			return true
		},
	},
	// At least 2 of the adjacent squares of a dragon must be empty.
	// If a dragon is at the edge or in the corner of the grid, the possible number of adjacent squares is reduced.
	{
		name: "survive",
		check: func(g *Grid, i int) bool {
			square := g.Squarei(i)
			if square == SquareDragon {
				possibleEmpty := g.NeighborCount(i, SquareEmpty, true, true)
				return possibleEmpty >= 2
			}
			return true
		},
	},
}

// Validate validates, if a grid is in a valid state.
// This applies all validation rules to the grid state.
// A valid state does not necessarily mean, that it can lead to a solution.
func Validate(g *Grid) bool {
	for i := range g.Squares {
		for _, rule := range rules {
			if !rule.check(g, i) {
				return false
			}
		}
	}
	return true
}

// ValidatePartial validates only some squares (ixs) of the whole grid.
// The goal is to optimize performance by validating only squares that were changed.
// This function is therefore best used when validating incrementally.
func ValidatePartial(g *Grid, ixs []int) bool {
	for _, i := range ixs {
		for _, rule := range rules {
			if !rule.check(g, i) {
				return false
			}
		}
	}
	return true
}
