package model

import (
	"fmt"
	"math/bits"
)

// Solve solves a world and returns the only possible solution or otherwise panics.
func Solve(w *World) *World {
	if !w.HasDistinctSolution() {
		fmt.Println("Solver: no distinct solution possible!")
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
		fmt.Println("-----")
	}
	return nil
}

// SolveSimple solves a world state only using the validation rules of the game.
func SolveSimple(w *World) *World {
	tries := 10
	if !w.HasDistinctSolution() {
		fmt.Println("SolveSimple: no distinct solution possible!")
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

			fmt.Println("Investigating square", i)

			// if only one option is possible for a sqaure.. go for it
			ops := knowledge.getOptions(cpy, i)
			if len(ops) == 1 {
				cpy.SetSquareByIndex(i, ops[0])
				fmt.Println("-> set the only possible option")
				fmt.Println(cpy)
				continue
			}

			// for all possible options, try and evaluate
			var ok []Square
			fmt.Println("-> evaluate options")
			for _, o := range ops {
				fmt.Printf("   -> option %c", squareSymbolsForCode[o])
				test := cpy.Clone()
				test.SetSquareByIndex(i, o)

				if !ValidateWorld(test) {
					// update knowledge
					fmt.Printf(" => NOT possible\n")
					knowledge.squareCannotBe(i, o)
					continue
				}

				// we compute all permutations that are possible when the neighbour squares are taken into account
				nis := test.GetNeighbourIndexes(i, false)
				permRes := knowledge.getPermutations(test, nis)
				fmt.Printf(" => permutations of %v: %v/%v", nis, permRes.valid, permRes.count)

				if permRes.valid == 0 {
					// update knowledge
					fmt.Println(" => [NOK]")
					knowledge.squareCannotBe(i, o)
				} else {
					fmt.Println(" => [ OK]")
					ok = append(ok, o)
				}
			}
			// if only one solution works.. use it
			if len(ok) == 1 {
				cpy.SetSquareByIndex(i, ok[0])
				fmt.Println("-> set the only possible option (after evaluating)")
				fmt.Println(cpy)
			}
		}
		if isSolved(cpy) {
			return cpy
		}

		fmt.Println("-----")
	}
	return nil
}

func isSolved(w *World) bool {
	return w.CountSquares(SquareUndefined) == 0 && ValidateWorld(w)
}

type solveRule (func(*World, int, *knowledge) *World)

var solveRules = []solveRule{
	// if a dragon is set, no neighbour square can be dragons
	func(w *World, i int, k *knowledge) *World {
		if w.GetSquareByIndex(i) == SquareDragon {
			k.squaresCannotBe(w.GetNeighbourIndexes(i, false), SquareDragon)
			fmt.Println("-> [ar] neighbour squares of dragon cannot be dragons")
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
				fmt.Println("-> [ar] fill adjacent squares to empty")
				fmt.Println(w)
			}
		}
		if w.GetSquareByIndex(i) == SquareUndefined {
			empty := w.CountAdjacentNeighbours(i, SquareEmpty)
			undef := w.CountAdjacentNeighbours(i, SquareUndefined)
			if empty+undef < 2 {
				k.squareCannotBe(i, SquareDragon)
				fmt.Println("-> [ar] square connot be a dragon because there could not be 2 empty adjacent square")
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
				fmt.Println("-> [ar] set dragons in neighbour squares if only one solution")
				fmt.Println(w)
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
				fmt.Println("-> [ar] set fire if more than 1 dragon around square")
				fmt.Println(w)
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
				fmt.Println("-> [ar] square cannot be fire if not at least 2 dragons could be around it")
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
				fmt.Println("-> [ar] if a empty square is set, there can maximum be one dragon around it")
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
				fmt.Println("-> [ar] set the only posible value of a square")
				fmt.Println(w)
			}
		}
		return w
	},
}
