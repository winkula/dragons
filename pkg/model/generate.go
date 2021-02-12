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
	g := best(c, func(g *Grid) float64 {
		return g.SolutionRating()
	})

	// we always normalize generated grids to prevent duplicates
	return g.Normalize()
}

func tryGenerate(width int, height int) *Grid {
	start := rand.Intn(width)
	fuzzyness := 4 * rand.Float64()
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
		work := g.Clone().SetDragon(index)
		if ValidateIncr(work, index, 2) {
			g = work
		}
	}

	g.FillUndefined(SquareAir)
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

	c := executeParallelLoop(duration, func(ctx context.Context, c chan<- *Grid) {
		c <- obfuscate(ctx, g, difficulty)
	})

	// take the best result
	return best(c, func(g *Grid) float64 {
		return g.PuzzleRating()
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
		// every puzzle must always have a distinct solution to be solvable
		return false
	}

	// TODO: optimize the difficulty check
	// stop early when difficulty is easy, dont check for medium then
	return CheckDifficulty(g, difficulty)
	//return difficulty == GetDifficulty(g)
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

func executeParallelLoop(timeout time.Duration, action func(ctx context.Context, c chan<- *Grid)) <-chan *Grid {
	return executeParallel(timeout, func(ctx context.Context, c chan<- *Grid) {
		for {
			action(ctx, c)
			if isTimeout(ctx) {
				break
			}
		}
	})
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

func best(c <-chan *Grid, evalFun func(*Grid) float64) *Grid {
	var best *Grid
	bestScore := -1000000000000.0
	for candidate := range c {
		score := evalFun(candidate)
		if score > bestScore {
			best = candidate
			bestScore = score
		}
	}
	return best
}

// TODO: improve
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
