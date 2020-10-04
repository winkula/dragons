package model

import (
	"math"
)

// Interestingness returns the interestingness of a grid.
// This is useful if we want to sort possible solved grids.
func (g *Grid) Interestingness() int {
	return int(float64(10.0*g.Density()) * g.Randomness())
}

// Density returns the density of a grid.
func (g *Grid) Density() (density int) {
	for _, val := range g.Squares {
		density += squareDensity[val]
	}
	return
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
