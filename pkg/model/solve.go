package model

import (
	"math/bits"
)

// Solve solves a puzzle only using the validation rules of the game.
// it returns the only possible solution or "nil" otherwise.
func Solve(g *Grid) *Grid {
	tries := 10
	if !IsDistinct(g) {
		debug("SolveSimple: no distinct solution possible!")
		return nil
	}

	cpy := g.Clone()
	knowledge := newKnowledge(cpy)
	for try := 0; try < tries; try++ {
		for i := range cpy.Squares {
			val := cpy.Squarei(i)
			if val != SquareUndefined {
				continue // square already defined
			}

			debug("Investigating square", i)

			// if only one option is possible for a sqaure.. go for it
			ops := knowledge.getOptions(cpy, i)
			if len(ops) == 1 {
				cpy.SetSquarei(i, ops[0])
				debug("-> set the only possible option")
				debug(cpy)
				continue
			}

			// for all possible options, try and evaluate
			ok := make([]Square, 0, len(options))
			debug("-> evaluate options")
			for _, o := range ops {
				//logd.Printf("   -> option %c", squareSymbolsForCode[o])
				test := cpy.Clone()
				test.SetSquarei(i, o)

				if !Validate(test) {
					// update knowledge
					//logd.Printf(" => NOT possible\n")
					knowledge.squareCannotBe(i, o)
					continue
				}

				// we compute all permutations that are possible when the neighbour squares are taken into account
				nis := test.NeighborIndicesi(i, false)
				permRes := knowledge.getPermutations(test, nis)
				//logd.Printf(" => permutations of %v: %v/%v", nis, permRes.valid, permRes.count)

				if permRes.valid == 0 {
					// update knowledge
					debug(" => [NOK]")
					knowledge.squareCannotBe(i, o)
				} else {
					debug(" => [ OK]")
					ok = append(ok, o)
				}
			}
			// if only one solution works.. use it
			if len(ok) == 1 {
				cpy.SetSquarei(i, ok[0])
				debug("-> set the only possible option (after evaluating)")
				debug(cpy)
			}
		}
		if isSolved(cpy) {
			return cpy
		}

		debug("-----")
	}
	return nil
}

// SolveDk solves a puzzle using domain knowledge.
// it returns the only possible solution or "nil" otherwise.
func SolveDk(g *Grid) *Grid {
	if !IsDistinct(g) {
		debug("Solver: no distinct solution possible!")
		return nil
	}
	cpy := g.Clone()
	knowledge := newKnowledge(cpy)
	for try := 0; try < 10; try++ {
		for i := range cpy.Squares {
			for _, r := range solveRules {
				r(cpy, i, knowledge)
				if isSolved(cpy) {
					return cpy
				}
			}
		}
		debug("-----")
	}
	return nil
}

// SolveBf solves a puzzle using a brute force strategy (enumerating all possible states).
func SolveBf(g *Grid) *Grid {
	solutions := Enumerate(g)
	if len(solutions) == 1 {
		return solutions[0]
	}
	return nil
}

func isSolved(g *Grid) bool {
	return g.CountSquares(SquareUndefined) == 0 && Validate(g)
}

type solveRule (func(*Grid, int, *knowledge) *Grid)

var solveRules = []solveRule{
	// if a dragon is set, no neighbour square can be dragons
	func(g *Grid, i int, k *knowledge) *Grid {
		if g.Squarei(i) == SquareDragon {
			k.squaresCannotBe(g.NeighborIndicesi(i, false), SquareDragon)
			debug("-> [ar] neighbour squares of dragon cannot be dragons")
		}
		return g
	},

	// if a dragon is set, at least two adjacent quares must be empty
	func(g *Grid, i int, k *knowledge) *Grid {
		if g.Squarei(i) == SquareDragon {
			empty := g.CountAdjacentNeighbours(i, SquareEmpty)
			undef := g.CountAdjacentNeighbours(i, SquareUndefined)
			if empty < 2 && empty+undef == 2 {
				for _, ni := range g.NeighborIndicesi(i, true) {
					if g.Squarei(ni) == SquareUndefined {
						g.SetSquarei(ni, SquareEmpty)
					}
				}
				debug("-> [ar] fill adjacent squares to empty")
				debug(g)
			}
		}
		if g.Squarei(i) == SquareUndefined {
			empty := g.CountAdjacentNeighbours(i, SquareEmpty)
			undef := g.CountAdjacentNeighbours(i, SquareUndefined)
			if empty+undef < 2 {
				k.squareCannotBe(i, SquareDragon)
				debug("-> [ar] square connot be a dragon because there could not be 2 empty adjacent square")
			}
		}
		return g
	},

	// if a fire square is set, there must be at least 2 dragons around it
	func(g *Grid, i int, k *knowledge) *Grid {
		if g.Squarei(i) == SquareFire {
			dragons := g.CountNeighbors(i, SquareDragon)
			undef := g.CountNeighbors(i, SquareUndefined)
			if dragons < 2 && dragons+undef == 2 {
				for _, ni := range g.NeighborIndicesi(i, false) {
					if g.Squarei(ni) == SquareUndefined {
						g.SetSquarei(ni, SquareDragon)
					}
				}
				debug("-> [ar] set dragons in neighbour squares if only one solution")
				debug(g)
			}
		}
		return g
	},

	// if a undefined square is surrounded by more than one dragon, there must be fire
	func(g *Grid, i int, k *knowledge) *Grid {
		if g.Squarei(i) == SquareUndefined {
			dragons := g.CountNeighbors(i, SquareDragon)
			if dragons > 1 {
				g.SetSquarei(i, SquareFire)
				debug("-> [ar] set fire if more than 1 dragon around square")
				debug(g)
			}
		}
		return g
	},

	// if a undefined square cannot be surrounded be at least 2 dragons, there can not be fire
	func(g *Grid, i int, k *knowledge) *Grid {
		if g.Squarei(i) == SquareUndefined {
			dragons := g.CountNeighbors(i, SquareDragon)
			undef := g.CountNeighbors(i, SquareUndefined)
			if dragons+undef < 2 {
				k.squareCannotBe(i, SquareFire)
				debug("-> [ar] square cannot be fire if not at least 2 dragons could be around it")
			}
		}
		return g
	},

	// if a empty square is set, there can maximum be one dragon around it
	func(g *Grid, i int, k *knowledge) *Grid {
		square := g.Squarei(i)
		if square == SquareEmpty {
			dragons := g.CountNeighbors(i, SquareDragon)
			if dragons == 1 {
				k.squaresCannotBe(g.NeighborIndicesi(i, false), SquareDragon)
				debug("-> [ar] if a empty square is set, there can maximum be one dragon around it")
			}
		}
		return g
	},

	// if a square is undefined but there is only one value possible (according to the knowledge db)
	// set the squares value
	func(g *Grid, i int, k *knowledge) *Grid {
		if g.Squarei(i) == SquareUndefined {
			if bits.OnesCount8(k.pv[i]) == 1 {
				val := Square(bits.TrailingZeros8(k.pv[i]))
				g.SetSquarei(i, val)
				debug("-> [ar] set the only posible value of a square")
				debug(g)
			}
		}
		return g
	},
}
