package model

import (
	"context"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// Generate generates a random but solved grid with the given dimensions.
// The duration parameter controls the time that is invested in generating
// the best or most interesting grid possible.
func Generate(width int, height int, duration time.Duration) *Grid {
	c := executeTimeBoundParallel(duration, func(ctx context.Context) *Grid {
		return generateBest(ctx, width, height)
	})

	// take the best result
	var best *Grid
	bestScore := -1
	for candidate := range c {
		score := candidate.Interestingness()
		if score > bestScore {
			best = candidate
			bestScore = score
		}
	}

	return best
}

func generateBest(ctx context.Context, width int, height int) *Grid {
	var best *Grid
	bestScore := -1
	for {
		if isTimeout(ctx) {
			break
		}
		candidate := generateRandomized(ctx, width, height)
		score := candidate.Interestingness()
		if score > bestScore {
			best = candidate
			bestScore = score
		}
	}
	return best
}

func generateInSequence(ctx context.Context, width int, height int) *Grid {
	chanceToSkip := 0.4
	g := New(width, height)
	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {
			i := g.Index(col, row)

			if g.Squarei(i) != SquareUndefined {
				continue // square already filled
			}

			if g.CountNeighbors(i, SquareDragon) > 0 {
				continue // already dragons in the neighbour squares
			}

			skipField := rand.Float64() <= chanceToSkip
			if skipField {
				continue // randomly skip field to make the generate algorithm non deterministic
			}

			// finally set the dragon
			g.SetDragon(i)
		}
	}
	return g
}

// TODO: rework this implementation!
func generateRandomized(ctx context.Context, width int, height int) *Grid {
	g := New(width, height).Fill(SquareEmpty)
	failsMax := 10000
	fails := 0
	size := g.Size()
	for fails < failsMax {
		if isTimeout(ctx) {
			break
		}
		i := rand.Intn(size)
		if g.Squarei(i) != SquareEmpty || g.CountNeighbors(i, SquareDragon) > 0 {
			fails++
			continue
		}
		suc := g.Clone().SetDragon(i)
		if !Validate(suc) {
			fails++
			continue
		}
		g = suc
	}
	return g
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

	c := executeTimeBoundParallel(duration, func(ctx context.Context) *Grid {
		return obfuscate(ctx, g, difficulty)
	})

	// take the best result
	var best *Grid
	bestScore := 0
	for candidate := range c {
		score := candidate.CountSquares(SquareUndefined)
		if score > bestScore {
			best = candidate
			bestScore = score
		}
	}
	return best
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
			index = (index + increment) % size
			if g.Squarei(index) != SquareUndefined {
				suc := g.Clone()
				suc.SetSquarei(index, SquareUndefined)
				if checkSolvable(suc, index, difficulty) {
					g = suc
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

func executeTimeBoundParallel(timeout time.Duration, action func(ctx context.Context) *Grid) chan *Grid {
	n := runtime.NumCPU()
	c := make(chan *Grid)

	var wg sync.WaitGroup
	wg.Add(n)

	// start go routines (one for each CPU)
	for i := 0; i < n; i++ {
		go func() {
			ctx, cancel := context.WithTimeout(context.TODO(), timeout)
			defer cancel() // releases resources
			c <- action(ctx)
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
