package model

import (
	"fmt"
	"strings"
)

// World represents the puzzles state.
type World struct {
	Width   int
	Height  int
	Squares []Square
}

// NewWorld creates a new world struct.
func NewWorld(width, height int) *World {
	return &World{
		Width:   width,
		Height:  height,
		Squares: make([]Square, width*height),
	}
}

func coordsToIndex(w *World, x int, y int) int {
	return w.Width*y + x
}

func coordsExist(w *World, x int, y int) bool {
	return x >= 0 && x < w.Width && y >= 0 && y < w.Height
}

// GetCoords gets the coordinates of a quare by the square index.
func (w *World) GetCoords(i int) (x, y int) {
	return i % w.Width, i / w.Width
}

// GetSquare gets the field value at coordinates x, y.
func (w *World) GetSquare(x int, y int) Square {
	if !coordsExist(w, x, y) {
		return SquareOut
	}
	return w.Squares[coordsToIndex(w, x, y)]
}

// SetSquare sets the squares value at coordinates x, y.
func (w *World) SetSquare(x int, y int, val Square) *World {
	if !coordsExist(w, x, y) {
		panic(fmt.Sprintf("Error: invalid coordinates (x = %v, y = %v, width = %v, height = %v)", x, y, w.Width, w.Height))
	}
	w.Squares[coordsToIndex(w, x, y)] = val
	return w
}

// GetNeighbours gets the neighbours field values.
func (w *World) GetNeighbours(x int, y int) []Square {
	return []Square{
		w.GetSquare(x-1, y-1),
		w.GetSquare(x, y-1),
		w.GetSquare(x+1, y-1),
		w.GetSquare(x-1, y),
		w.GetSquare(x+1, y),
		w.GetSquare(x-1, y+1),
		w.GetSquare(x, y+1),
		w.GetSquare(x+1, y+1),
	}
}

// Clone creates an exact deep copy of a world.
func (w *World) Clone() *World {
	worldCopy := *w
	worldCopy.Squares = make([]Square, len(w.Squares))
	copy(worldCopy.Squares, w.Squares)
	return &worldCopy
}

func (w *World) String() string {
	var sb strings.Builder
	for i, val := range w.Squares {

		sb.WriteRune(getSymbol(val))
		sb.WriteRune(' ')

		if i%w.Width == w.Width-1 && i != len(w.Squares)-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func getSymbol(val Square) rune {
	return squareSymbols[val]
}
