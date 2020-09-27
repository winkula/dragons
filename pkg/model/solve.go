package model

import (
	"math"
	"math/bits"
)

// Solve solves a puzzle only using the validation rules of the game.
// it returns the only possible solution or "nil" otherwise.
func Solve(g *Grid) *Grid {
	return SolveHuman(g, DifficultyHard)
}

// SolveHuman solves a puzzle but only if it is easier than a given difficulty level.
func SolveHuman(g *Grid, difficulty Difficulty) *Grid {
	if !IsDistinct(g) {
		return nil // no distinct solution exists
	}

	maxPerm := maxPermCount(difficulty)

	work := g.Clone()
	k := newKnowledge(work)
	dirty := true // as long as this flag is set, we need to check every square of the grid
	for dirty {
		dirty = false
		for i := range work.Squares {
			val := work.Squarei(i)
			if val != SquareUndefined {
				continue // square already defined
			}

			debug(" ")
			debug(" ")
			debug(" ")
			debug("========== Investigating square", i, "==========")
			debug(Render(work, k, i))

			// if only one option is possible for a square.. go for it
			ops := k.getOptions(work, i)
			// if len(ops) == 1 {
			// 	work.SetSquarei(i, ops[0])
			// 	dirty = true
			// 	k.squareIs(i, ops[0])
			// 	// we don't need to validate here as it's the only possible option anyway
			// 	// IMPORTANT: this is only correct if we know, that a solution exists!
			// 	debug("-> set the only possible option")
			// 	debug(Render(work, k, i))
			// 	continue
			// }

			// for all possible options, try and evaluate
			ok := make([]Square, 0, len(options))
			debug("-> evaluate options")
			for _, o := range ops {

				// if the grid is already invalid after setting the square value,
				// we can immediatelty go to the next option and can save us the permutation check
				test := work.Clone()
				valid := test.SetSquareiAndValidate(i, o)
				if !valid {
					// update knowledge
					debug("   -> option", string(squareSymbols[o]), "[NOK] (invalid state)")
					dirty = true
					k.squareCannotBe(i, o)
					continue
				}

				// we compute all permutations that are possible when the neighbour squares are taken into account
				nis := test.NeighborIndicesi(i, false)
				permCount := k.getPermCount(test, nis)

				// too much permutations...
				// this algorithm should simulate the human that solves the puzzle
				// with the maxPermutationsToEvaluate parameter, we can fine tune how many permutations the human can/will evaluate
				if permCount > maxPerm {
					debug("   -> option", string(squareSymbols[o]), "[ OK] (max permutations is too big for the difficulty level, so we can't exclude this option)")
					debug("      permCount:", permCount, "maxPerm:", maxPerm)
					ok = append(ok, o)
					continue
				}

				permRes := k.getPermutations(test, nis)
				if permRes.valid == 0 {
					// not a valid option for this square: update knowledge
					debug("   -> option", string(squareSymbols[o]), "[NOK] (no valid permutations)")
					debug("      permutations of", nis, "valid:", permRes.valid, "total:", permRes.total)
					k.squareCannotBe(i, o)
					dirty = true
					continue
				}

				// at least one valid permutation exists

				debug("   -> option", string(squareSymbols[o]), "[ OK]")
				debug("      permutations of", nis, "valid:", permRes.valid, "total:", permRes.total)
				ok = append(ok, o)
			}
			// if only one solution works... use it
			// if not: try the next square...
			if len(ok) == 1 {
				work.SetSquarei(i, ok[0])
				dirty = true
				k.squareIs(i, ok[0])
				debug("-> set the only possible option (after evaluating)")
				debug("\n" + Render(work, k, i))
			}
		}

		// after going through all squares, we check if the puzzle is already solved
		// if not, continue until we reach the tries timeout
		if isSolved(work) {
			return work
		}
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

// GetDifficulty gets the difficulty of a puzzle.
func GetDifficulty(g *Grid) Difficulty {
	for _, difficulty := range []Difficulty{DifficultyEasy, DifficultyMedium} {
		solution := SolveHuman(g, difficulty)
		if solution != nil {
			return difficulty
		}
	}
	return DifficultyHard
}

func isSolved(g *Grid) bool {
	if g == nil {
		return false
	}
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
			empty := g.NeighborCount4(i, SquareEmpty)
			undef := g.NeighborCount4(i, SquareUndefined)
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
			empty := g.NeighborCount4(i, SquareEmpty)
			undef := g.NeighborCount4(i, SquareUndefined)
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
			dragons := g.NeighborCount8(i, SquareDragon)
			undef := g.NeighborCount8(i, SquareUndefined)
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
			dragons := g.NeighborCount8(i, SquareDragon)
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
			dragons := g.NeighborCount8(i, SquareDragon)
			undef := g.NeighborCount8(i, SquareUndefined)
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
			dragons := g.NeighborCount8(i, SquareDragon)
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

func maxPermCount(difficulty Difficulty) int {
	switch difficulty {
	case DifficultyEasy:
		return 1 // only one valid possibilities
	case DifficultyMedium:
		return 9 // 3^2 (all possibilities for 2 fields)
	default:
		return math.MaxUint32
	}
}
