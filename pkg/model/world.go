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

func indexExists(w *World, index int) bool {
	return index >= 0 && index < w.Width*w.Height
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

// GetSquareByIndex gets the field value at the specified index.
func (w *World) GetSquareByIndex(index int) Square {
	if !indexExists(w, index) {
		panic(fmt.Sprintf("Error: invalid index (index = %v, width = %v, height = %v)", index, w.Width, w.Height))
	}
	return w.Squares[index]
}

// SetSquare sets the squares value at coordinates x, y.
func (w *World) SetSquare(x int, y int, val Square) *World {
	if !coordsExist(w, x, y) {
		panic(fmt.Sprintf("Error: invalid coordinates (x = %v, y = %v, width = %v, height = %v)", x, y, w.Width, w.Height))
	}
	w.Squares[coordsToIndex(w, x, y)] = val
	return w
}

// SetSquareByIndex sets the squares value at the specified index.
func (w *World) SetSquareByIndex(index int, val Square) *World {
	if !indexExists(w, index) {
		panic(fmt.Sprintf("Error: invalid index (index = %v, width = %v, height = %v)", index, w.Width, w.Height))
	}
	w.Squares[index] = val
	return w
}

// FillSquares fills all squares.
func (w *World) FillSquares(val Square) *World {
	for i := range w.Squares {
		w.Squares[i] = val
	}
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

// GetNeighbourIndexes gets the neighbours indexes.
func (w *World) GetNeighbourIndexes(i int) []int {
	x, y := w.GetCoords(i)
	return []int{
		coordsToIndex(w, x-1, y-1),
		coordsToIndex(w, x, y-1),
		coordsToIndex(w, x+1, y-1),
		coordsToIndex(w, x-1, y),
		coordsToIndex(w, x+1, y),
		coordsToIndex(w, x-1, y+1),
		coordsToIndex(w, x, y+1),
		coordsToIndex(w, x+1, y+1),
	}
}

// GetAdjacentNeighbourIndexes gets the adjacent neighbours indexes.
func (w *World) GetAdjacentNeighbourIndexes(i int) []int {
	x, y := w.GetCoords(i)
	return []int{
		coordsToIndex(w, x, y-1),
		coordsToIndex(w, x-1, y),
		coordsToIndex(w, x+1, y),
		coordsToIndex(w, x, y+1),
	}
}

// GetAdjacentNeighbours gets the adjacent neighbours field values.
func (w *World) GetAdjacentNeighbours(x int, y int) []Square {
	return []Square{
		w.GetSquare(x, y-1),
		w.GetSquare(x-1, y),
		w.GetSquare(x+1, y),
		w.GetSquare(x, y+1),
	}
}

// CountNeighbours counts the neighboured squares that match the given type.
func (w *World) CountNeighbours(index int, square Square) int {
	x, y := w.GetCoords(index)
	ns := w.GetNeighbours(x, y)
	count := 0
	for _, v := range ns {
		if v == square {
			count++
		}
	}
	return count
}

// CountAdjacentNeighbours counts the adacent neighboured squares that match the given type.
func (w *World) CountAdjacentNeighbours(index int, square Square) int {
	x, y := w.GetCoords(index)
	ns := w.GetAdjacentNeighbours(x, y)
	count := 0
	for _, v := range ns {
		if v == square {
			count++
		}
	}
	return count
}

// Clone creates an exact deep copy of a world.
func (w *World) Clone() *World {
	worldCopy := *w
	worldCopy.Squares = make([]Square, len(w.Squares))
	copy(worldCopy.Squares, w.Squares)
	return &worldCopy
}

// SetDragon sets a dragon to a specific square and computes where fire must be.
func (w *World) SetDragon(index int) {
	x, y := w.GetCoords(index)
	w.SetSquare(x, y, SquareDragon)
	ns := w.GetNeighbours(x, y)
	for i := range ns {
		numDragons := w.CountNeighbours(i, SquareDragon)
		if numDragons > 1 {
			w.SetSquareByIndex(i, SquareFire)
		}
	}
}

// Interestingness returns the interestingness of a world.
// This is useful if we want to sort possible solved worlds.
func (w *World) Interestingness() (interestingness int) {
	for _, val := range w.Squares {
		interestingness += squareInterestingness[val]
	}
	return interestingness
}

// Size returns a worlds number of squares.
func (w *World) Size() int {
	return w.Width * w.Height
}

// CountSquares fills all squares.
func (w *World) CountSquares(needle Square) int {
	return countSquares(w.Squares, needle)
}

func (w *World) String() string {
	var sb strings.Builder
	sb.WriteString("┌")
	if w.Width > 1 {
		pad := (2*w.Width + 3 - 7) / 2
		sb.WriteString(strings.Repeat("─", pad))
		sb.WriteString(fmt.Sprintf("╴%vx%v╶", w.Width, w.Height))
		sb.WriteString(strings.Repeat("─", pad))
	} else {
		sb.WriteString(strings.Repeat("─", 2*w.Width+1))
	}
	sb.WriteString("┐\n")
	for i, val := range w.Squares {
		if i%w.Width == 0 {
			sb.WriteString("│ ")
		}
		sb.WriteRune(getSymbol(val))
		sb.WriteRune(' ')
		if i%w.Width == w.Width-1 {
			sb.WriteString("│\n")
		}
	}
	sb.WriteString("└")
	sb.WriteString(strings.Repeat("─", 2*w.Width+1))
	sb.WriteString("┘")
	sb.WriteString("\nCode: ")
	for i, val := range w.Squares {
		sb.WriteRune(getSymbolForCode(val))
		if i%w.Width == w.Width-1 && i < w.Width*w.Height-1 {
			sb.WriteString(",")
		}
	}
	return sb.String()
}

func getSymbol(val Square) rune {
	return squareSymbols[val]
}

func getSymbolForCode(val Square) rune {
	return squareSymbolsForCode[val]
}

func countSquares(squares []Square, needle Square) (count int) {
	for _, s := range squares {
		if s == needle {
			count++
		}
	}
	return
}
