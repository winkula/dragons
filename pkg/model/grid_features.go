package model

import (
	"math"
	"math/big"
)

// SolutionRating returns the interestingness of a grid.
// This is useful if we want to sort possible solved grids.
func (g *Grid) SolutionRating() float64 {
	return (g.Density() + g.Randomness()) / 2.0
}

// Density returns the density of a grid.
func (g *Grid) Density() float64 {
	filled := 0
	for _, val := range g.Squares {
		filled += val.Density()
	}
	return float64(filled) / float64(g.Size())
}

// Randomness returns the randomness of a grid.
func (g *Grid) Randomness() float64 {
	cols := 0.0
	rows := 0.0
	for x := 0; x < g.Width; x++ {
		for y := 0; y < g.Height; y++ {
			val := g.Square(x, y)

			incx := 1.0 / math.Min(float64(x+1), 3.0)
			if val == g.Square(x-1, y) {
				rows += incx
			}
			if val == g.Square(x-2, y) {
				rows += incx
			}
			if val == g.Square(x-3, y) {
				rows += incx
			}

			incy := 1.0 / math.Min(float64(y+1), 3.0)
			if val == g.Square(x, y-1) {
				cols += incy
			}
			if val == g.Square(x, y-2) {
				cols += incy
			}
			if val == g.Square(x, y-3) {
				cols += incy
			}
		}
	}

	return 1.0 - (cols+rows)/(float64(g.Size()))/2.0
}

// Undefinedness returns the undefinedness of a unsolved puzzle grid.
// This is especially useful to find "interesting" puzzle.
func (g *Grid) Undefinedness() float64 {
	undef := g.CountSquares(SquareUndefined)
	all := g.Size()
	return float64(undef) / float64(all)
}

// PuzzleRating is used to generate puzzles from solutions.
// This value is used to find the best obfuscated puzzle among others.
func (g *Grid) PuzzleRating() float64 {
	return GetAvgOptions(g, DifficultyEasy)
}

// ID creates a unique identifier for a puzzle.
// This helps when we want to normalize puzzles.
//
// Code is constructed as follows:
//
// most significant bit <------------------------------- least significant bit
// [ last square ] ... [ square 0 (2bits) ][ height (5bits) ][ width (5bits) ]
func (g *Grid) ID() *big.Int {
	code := big.NewInt(0)

	code.Or(code, big.NewInt(0).Lsh(big.NewInt(int64(g.Width)), uint(0)))
	code.Or(code, big.NewInt(0).Lsh(big.NewInt(int64(g.Height)), uint(5)))

	for i, val := range g.Squares {
		squareValue := big.NewInt(int64(val))
		s := big.NewInt(0).Lsh(squareValue, uint(10+2*i))
		code.Or(code, s)
	}
	return code
}
