package model

import (
	"strings"
)

const outsideFieldValue = -1

var m = map[int]rune{
	0: '.',
	1: 'X',
	2: 'F',
	3: 'D',
}

// World represents the puzzles state.
type World struct {
	Width  int
	Height int

	fields []int
}

// NewWorld creates a new world struct.
func NewWorld(width, height int) *World {
	return &World{
		Width:  width,
		Height: height,
		fields: make([]int, width*height),
	}
}

func coordsToIndex(w *World, x int, y int) int {
	return w.Width*y + x
}

func coordsExist(w *World, x int, y int) bool {
	return x >= 0 && x < w.Width && y >= 0 && y < w.Height
}

// GetField gets the field value at coordinates x, y.
func (w *World) GetField(x int, y int) int {
	if !coordsExist(w, x, y) {
		return outsideFieldValue
	}
	return w.fields[coordsToIndex(w, x, y)]
}

// SetField sets the fields value at coordinates x, y.
func (w *World) SetField(x int, y int, val int) {
	if !coordsExist(w, x, y) {
		panic("Error: invalid coordinates")
	}
	w.fields[coordsToIndex(w, x, y)] = val
}

// GetNeighbours gets the neighbours field values.
func (w *World) GetNeighbours(x int, y int) []int {
	return []int{
		w.GetField(x-1, y-1),
		w.GetField(x, y-1),
		w.GetField(x+1, y-1),
		w.GetField(x-1, y),
		w.GetField(x+1, y),
		w.GetField(x-1, y+1),
		w.GetField(x, y+1),
		w.GetField(x+1, y+1),
	}
}

func (w *World) String() string {
	var sb strings.Builder
	for i, val := range w.fields {

		sb.WriteRune(getSymbol(val))
		sb.WriteRune(' ')

		if i%w.Width == w.Width-1 && i != len(w.fields)-1 {
			sb.WriteRune('\n')
		}
	}
	return sb.String()
}

func getSymbol(val int) rune {
	return m[val]
}
