package model

import (
	"math"
)

// Solve solves a puzzle using multiple algorithms if necessary.
// it returns the only possible solution or "nil" otherwise.
func Solve(g *Grid) *Grid {
	solution := SolveIterative(g, DifficultyHard)
	if solution != nil {
		return solution
	}

	// brute force algorithm as fallback
	return SolveBruteForce(g)
}

// SolveIterative solves a puzzle by iteratively checking possible options for a given square.
// It only uses the validation rules of the game - no domain knowledge required.
// Can be tuned by the maximum number of possible permutations that can exist in order to exclude a option.
// It returns the only possible solution or "nil" otherwise.
func SolveIterative(g *Grid, difficulty Difficulty) *Grid {
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
			debugGrid(work, i)

			// if only one option is possible for a square.. go for it
			ops := k.getOptions(work, i)
			// if len(ops) == 1 {
			// 	work.SetSquarei(i, ops[0])
			// 	dirty = true
			// 	k.squareIs(i, ops[0])
			// 	// we don't need to validate here as it's the only possible option anyway
			// 	// IMPORTANT: this is only correct if we know, that a solution exists!
			// 	debug("-> set the only possible option")
			// 	debugGrid(work, i)
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
					debug("   -> option", string(o.Symbol()), "[NOK] (invalid state)")
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
					debug("   -> option", string(o.Symbol()), "[ OK] (max permutations is too big for the difficulty level, so we can't exclude this option)")
					debug("      permCount:", permCount, "maxPerm:", maxPerm)
					ok = append(ok, o)
					continue
				}

				permRes := k.getPermutations(test, nis)
				if permRes.valid == 0 {
					// not a valid option for this square: update knowledge
					debug("   -> option", string(o.Symbol()), "[NOK] (no valid permutations)")
					debug("      permutations of", nis, "valid:", permRes.valid, "total:", permRes.total)
					k.squareCannotBe(i, o)
					dirty = true
					continue
				}

				// at least one valid permutation exists

				debug("   -> option", string(o.Symbol()), "[ OK]")
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
				debug("\n" + Render(work, i))
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

// SolveTechnically solves a puzzle applying fix techniques (using domain knowledge).
// It returns the only possible solution or "nil" otherwise.
func SolveTechnically(g *Grid, difficulty Difficulty) *Grid {
	if !IsDistinct(g) {
		debug("Solver: no distinct solution possible!")
		return nil
	}
	cpy := g.Clone()
	knowledge := newKnowledge(cpy)

	dirty := true // as long as this flag is set, we need to check every square of the grid
	for dirty {
		dirty = false
		for i := range cpy.Squares {
			for _, rule := range solveTechniques[difficulty] {
				if rule(cpy, i, knowledge) == solveResultFillSquare {
					// set dirty flag if any rule could be applied
					dirty = true
				}
				if isSolved(cpy) {
					return cpy
				}
			}
		}
	}
	debug("-----")

	return nil
}

// SolveBruteForce solves a puzzle using a brute force strategy (enumerating all possible states).
// It's the only algorithm that finds solutions to the hardest puzzles.
// It returns the only possible solution or "nil" otherwise.
func SolveBruteForce(g *Grid) *Grid {
	solutions := EnumerateLimited(g)
	if len(solutions) == 1 {
		return solutions[0]
	}
	return nil
}

// GetDifficulty gets the difficulty of a puzzle.
func GetDifficulty(g *Grid) Difficulty {
	
	if SolveTechnically(g, DifficultyEasy) != nil {
		if g.PuzzleRating() >= 0.35 {
			// easy puzzles must be solvable with applying easy techniques only (domain knowledge)
			return DifficultyEasy
		}		
		return DifficultyMedium
	}

	if SolveIterative(g, DifficultyEasy) != nil {
		// medium puzzles must be solvable using the SolveIterative algorithm with parameter "easy"
		// it should not be solvable with "SolveTechnically"
		return DifficultyHard
	}

	// hard puzzles have no restriction in being solvable using a specific algorithm
	// sometimes, brute force is the only option to solve a "hard" puzzle
	return DifficultyBrutal
}

// GetAvgOptions gets the average options of possible solve techniques that can be applied per square.
func GetAvgOptions(g *Grid, difficulty Difficulty) float64 {
	count := 0
	for i := range g.Squares {
		for _, rule := range solveTechniques[difficulty] {
			cpy := g.Clone()
			knowledge := newKnowledge(cpy)
			count += int(rule(cpy, i, knowledge))
		}
	}
	return float64(count) / float64(g.Size())
}

func anyRuleApplies(g *Grid, i int) bool {
	cpy := g.Clone()
	knowledge := newKnowledge(cpy)
	for _, rule := range solveRulesEasy {
		if rule(cpy, i, knowledge) > solveResultNone {
			return true
		}
	}
	return false
}

func isSolved(g *Grid) bool {
	if g == nil {
		return false
	}
	return g.CountSquares(SquareUndefined) == 0 && Validate(g)
}

func maxPermCount(difficulty Difficulty) int {
	switch difficulty {
	case DifficultyEasy:
		return 1 // only one valid possibilities
	case DifficultyMedium:
		return 3 // 3^1 (all possibilities for 1 field)
	case DifficultyHard:
		return 9 // 3^2 (all possibilities for 2 fields)
	default:
		return math.MaxUint32
	}
}
