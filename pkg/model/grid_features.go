package model

import (
	"math"
)

// Interestingness returns the interestingness of a grid.
// This is useful if we want to sort possible solved grids.
func (g *Grid) Interestingness() float64 {
	return (g.Density() + g.Randomness()) / 2.0
}

// Density returns the density of a grid.
func (g *Grid) Density() float64 {
	filled := 0
	for _, val := range g.Squares {
		filled += squareDensity[val]
	}
	g.Size()
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
