package model

import (
	"context"
	"math"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Generate generates a random but solved grid with the given dimensions.
// The duration parameter controls the time that is invested in generating
// the best or most interesting grid possible.
func Generate(width int, height int, duration time.Duration) *Grid {
	c := executeParallel(duration, func(ctx context.Context, c chan<- *Grid) {
		for {
			if isTimeout(ctx) {
				break
			}
			g := tryGenerate(width, height)
			if g != nil {
				c <- g
			}
		}
	})

	// take the most interesting result
	return best(c, func(g *Grid) int {
		return g.Interestingness()
	})
}

func tryGenerate(width int, height int) *Grid {
	start := rand.Intn(width)
	fuzzyness := 3 * rand.Float64()
	g := New(width, height)
	size := g.Size()
	for i := start; i < size; i++ {
		index := fuzzyIndex(i, size, height, fuzzyness)

		if g.Squarei(index) != SquareUndefined {
			continue // square already filled
		}
		if g.NeighborCount8(index, SquareDragon) > 0 {
			continue // already dragons in the neighbour squares
		}

		// set the dragon and validate the grid
		// TODO: use ValidatePartial??
		work := g.Clone().SetDragon(index)
		if Validate(work) {
			g = work
		}
	}

	g.FillUndefined(SquareEmpty)
	if Validate(g) {
		return g
	}

	return nil // so valid solution found
}

// Obfuscate creates a puzzle from a given solved or partially solved puzzle.
// The difficulty parameter controls how hard the puzzle would be to solve for a human.
//
// The algorithm works like this:
// It takes a grid state and incrementally sets random squares to "undefined".
// After every step, it verifies if the puzzle is still solvable (i.e. has a distinct solution).
func Obfuscate(g *Grid, difficulty Difficulty, duration time.Duration) *Grid {
	if g.IsUndefined() {
		panic("generateInternal: all squares undefined")
	}

	c := executeParallel(duration, func(ctx context.Context, c chan<- *Grid) {
		c <- obfuscate(ctx, g, difficulty)
	})

	// take the best result
	return best(c, func(g *Grid) int {
		return g.CountSquares(SquareUndefined)
	})
}

func obfuscate(ctx context.Context, g *Grid, difficulty Difficulty) *Grid {
	size := g.Size()
	seed := rand.Intn(size)
	increment := 997 // must be a prime number, see: https://en.wikipedia.org/wiki/Full_cycle
	for try := 0; try < size; try++ {
		if isTimeout(ctx) {
			break
		}
		index := seed + try
		for i := 0; i < size; i++ {
			index = (index + increment) % size // next prng value
			if g.Squarei(index) != SquareUndefined {
				work := g.Clone()
				work.SetSquarei(index, SquareUndefined)
				if checkSolvable(work, index, difficulty) {
					g = work
					break
				}
			}
		}
	}
	return g
}

func checkSolvable(g *Grid, index int, difficulty Difficulty) bool {
	if !IsDistinct(g) {
		return false
	}

	if difficulty == DifficultyHard {
		return true
	}

	solution := SolveHuman(g, difficulty)
	if solution == nil {
		return false
	}

	return true
}

func executeParallel(timeout time.Duration, action func(ctx context.Context, c chan<- *Grid)) <-chan *Grid {
	n := runtime.NumCPU()
	c := make(chan *Grid)

	var wg sync.WaitGroup
	wg.Add(n)

	// start go routines (one for each CPU)
	for i := 0; i < n; i++ {
		go func() {
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel() // releases resources
			action(ctx, c)
			wg.Done()
		}()
	}

	// close the channel when all go routines are finish
	go func() {
		wg.Wait()
		close(c)
	}()

	return c
}

func isTimeout(ctx context.Context) bool {
	// check context to see if we should terminate
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func best(c <-chan *Grid, evalFun func(*Grid) int) *Grid {
	var best *Grid
	bestScore := -1
	for candidate := range c {
		score := evalFun(candidate)
		if score > bestScore {
			best = candidate
			bestScore = score
		}
	}
	return best
}

func fuzzyIndex(i int, size int, height int, fuzzyness float64) int {
	index := i

	index += int(math.Round(rand.NormFloat64()*fuzzyness)) * height // skip rows
	index += int(math.Round(rand.NormFloat64() * fuzzyness))        // skip columns

	if index < 0 {
		return 0
	}
	if index >= size {
		return size - 1
	}

	return index
}
