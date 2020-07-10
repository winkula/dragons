package model

import (
	"math/bits"
)

// Solve solves a puzzle only using the validation rules of the game.
// it returns the only possible solution or "nil" otherwise.
func Solve(w *World) *World {
	tries := 10
	if !HasDistinctSolution(w) {
		logd.Println("SolveSimple: no distinct solution possible!")
		return nil
	}

	cpy := w.Clone()
	knowledge := newKnowledge(cpy)
	for try := 0; try < tries; try++ {
		for i := range cpy.Squares {
			val := cpy.GetSquareByIndex(i)
			if val != SquareUndefined {
				continue // square already defined
			}

			logd.Println("Investigating square", i)

			// if only one option is possible for a sqaure.. go for it
			ops := knowledge.getOptions(cpy, i)
			if len(ops) == 1 {
				cpy.SetSquareByIndex(i, ops[0])
				logd.Println("-> set the only possible option")
				logd.Println(cpy)
				continue
			}

			// for all possible options, try and evaluate
			var ok []Square
			logd.Println("-> evaluate options")
			for _, o := range ops {
				logd.Printf("   -> option %c", squareSymbolsForCode[o])
				test := cpy.Clone()
				test.SetSquareByIndex(i, o)

				if !Validate(test) {
					// update knowledge
					logd.Printf(" => NOT possible\n")
					knowledge.squareCannotBe(i, o)
					continue
				}

				// we compute all permutations that are possible when the neighbour squares are taken into account
				nis := test.GetNeighbourIndexes(i, false)
				permRes := knowledge.getPermutations(test, nis)
				logd.Printf(" => permutations of %v: %v/%v", nis, permRes.valid, permRes.count)

				if permRes.valid == 0 {
					// update knowledge
					logd.Println(" => [NOK]")
					knowledge.squareCannotBe(i, o)
				} else {
					logd.Println(" => [ OK]")
					ok = append(ok, o)
				}
			}
			// if only one solution works.. use it
			if len(ok) == 1 {
				cpy.SetSquareByIndex(i, ok[0])
				logd.Println("-> set the only possible option (after evaluating)")
				logd.Println(cpy)
			}
		}
		if isSolved(cpy) {
			return cpy
		}

		logd.Println("-----")
	}
	return nil
}

// SolveDk solves a puzzle using domain knowledge.
// it returns the only possible solution or "nil" otherwise.
func SolveDk(w *World) *World {
	if !HasDistinctSolution(w) {
		logd.Println("Solver: no distinct solution possible!")
		return nil
	}
	cpy := w.Clone()
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
		logd.Println("-----")
	}
	return nil
}

// SolveBf solves a puzzle using a brute force strategy (enumerating all possible states).
func SolveBf(w *World) *World {
	solutions := Enumerate(w)
	if len(solutions) == 1 {
		return solutions[0]
	}
	return nil
}

func isSolved(w *World) bool {
	return w.CountSquares(SquareUndefined) == 0 && Validate(w)
}

type solveRule (func(*World, int, *knowledge) *World)

var solveRules = []solveRule{
	// if a dragon is set, no neighbour square can be dragons
	func(w *World, i int, k *knowledge) *World {
		if w.GetSquareByIndex(i) == SquareDragon {
			k.squaresCannotBe(w.GetNeighbourIndexes(i, false), SquareDragon)
			logd.Println("-> [ar] neighbour squares of dragon cannot be dragons")
		}
		return w
	},

	// if a dragon is set, at least two adjacent quares must be empty
	func(w *World, i int, k *knowledge) *World {
		if w.GetSquareByIndex(i) == SquareDragon {
			empty := w.CountAdjacentNeighbours(i, SquareEmpty)
			undef := w.CountAdjacentNeighbours(i, SquareUndefined)
			if empty < 2 && empty+undef == 2 {
				for _, ni := range w.GetNeighbourIndexes(i, true) {
					if w.GetSquareByIndex(ni) == SquareUndefined {
						w.SetSquareByIndex(ni, SquareEmpty)
					}
				}
				logd.Println("-> [ar] fill adjacent squares to empty")
				logd.Println(w)
			}
		}
		if w.GetSquareByIndex(i) == SquareUndefined {
			empty := w.CountAdjacentNeighbours(i, SquareEmpty)
			undef := w.CountAdjacentNeighbours(i, SquareUndefined)
			if empty+undef < 2 {
				k.squareCannotBe(i, SquareDragon)
				logd.Println("-> [ar] square connot be a dragon because there could not be 2 empty adjacent square")
			}
		}
		return w
	},

	// if a fire square is set, there must be at least 2 dragons around it
	func(w *World, i int, k *knowledge) *World {
		if w.GetSquareByIndex(i) == SquareFire {
			dragons := w.CountNeighbours(i, SquareDragon)
			undef := w.CountNeighbours(i, SquareUndefined)
			if dragons < 2 && dragons+undef == 2 {
				for _, ni := range w.GetNeighbourIndexes(i, false) {
					if w.GetSquareByIndex(ni) == SquareUndefined {
						w.SetSquareByIndex(ni, SquareDragon)
					}
				}
				logd.Println("-> [ar] set dragons in neighbour squares if only one solution")
				logd.Println(w)
			}
		}
		return w
	},

	// if a undefined square is surrounded by more than one dragon, there must be fire
	func(w *World, i int, k *knowledge) *World {
		if w.GetSquareByIndex(i) == SquareUndefined {
			dragons := w.CountNeighbours(i, SquareDragon)
			if dragons > 1 {
				w.SetSquareByIndex(i, SquareFire)
				logd.Println("-> [ar] set fire if more than 1 dragon around square")
				logd.Println(w)
			}
		}
		return w
	},

	// if a undefined square cannot be surrounded be at least 2 dragons, there can not be fire
	func(w *World, i int, k *knowledge) *World {
		if w.GetSquareByIndex(i) == SquareUndefined {
			dragons := w.CountNeighbours(i, SquareDragon)
			undef := w.CountNeighbours(i, SquareUndefined)
			if dragons+undef < 2 {
				k.squareCannotBe(i, SquareFire)
				logd.Println("-> [ar] square cannot be fire if not at least 2 dragons could be around it")
			}
		}
		return w
	},

	// if a empty square is set, there can maximum be one dragon around it
	func(w *World, i int, k *knowledge) *World {
		square := w.GetSquareByIndex(i)
		if square == SquareEmpty {
			dragons := w.CountNeighbours(i, SquareDragon)
			if dragons == 1 {
				k.squaresCannotBe(w.GetNeighbourIndexes(i, false), SquareDragon)
				logd.Println("-> [ar] if a empty square is set, there can maximum be one dragon around it")
			}
		}
		return w
	},

	// if a square is undefined but there is only one value possible (according to the knowledge db)
	// set the squares value
	func(w *World, i int, k *knowledge) *World {
		if w.GetSquareByIndex(i) == SquareUndefined {
			if bits.OnesCount8(k.pv[i]) == 1 {
				val := Square(bits.TrailingZeros8(k.pv[i]))
				w.SetSquareByIndex(i, val)
				logd.Println("-> [ar] set the only posible value of a square")
				logd.Println(w)
			}
		}
		return w
	},
}
