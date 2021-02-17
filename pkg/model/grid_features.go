package model

import (
	"math/big"
)

// SolutionRating returns the interestingness of a grid.
// This is useful if we want to sort possible solved grids.
// Density returns the density of a grid (number of dragons percentually).
func (g *Grid) SolutionRating() float64 {
	return float64(g.CountSquares(SquareDragon)) / float64(g.Size())
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
//
// TODO: use a sigmoid function or similar to clamp the value always between 0.0 and 1.0
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
