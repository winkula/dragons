package model

import (
	"fmt"
	"strings"
)

const maxNumNeighbors = 8

var neighbours = []struct {
	x        int
	y        int
	adjacent bool
}{
	{x: -1, y: -1, adjacent: false},
	{x: 0, y: -1, adjacent: true},
	{x: 1, y: -1, adjacent: false},
	{x: -1, y: 0, adjacent: true},
	//{x: 0, y: 0, adjacent: false}, // self
	{x: 1, y: 0, adjacent: true},
	{x: -1, y: 1, adjacent: false},
	{x: 0, y: 1, adjacent: true},
	{x: 1, y: 1, adjacent: false},
}

// Grid represents the puzzles state.
type Grid struct {
	Width   int
	Height  int
	Squares []Square
}

// New creates a new grid struct.
func New(width int, height int) *Grid {
	return &Grid{
		Width:   width,
		Height:  height,
		Squares: make([]Square, width*height),
	}
}

// Coords gets the coordinates of a quare given the squares index.
func (g *Grid) Coords(i int) (x int, y int) {
	return i % g.Width, i / g.Width
}

// Index returns the square index from its coordinates.
func (g *Grid) Index(x int, y int) (int, bool) {
	if !coordsExist(g, x, y) {
		return -1, false // index does not exist
	}
	i := coordsToIndex(g, x, y)
	return i, true
}

// Square gets the field value at coordinates x, y.
func (g *Grid) Square(x int, y int) Square {
	if !coordsExist(g, x, y) {
		return SquareOut
	}
	return g.Squares[coordsToIndex(g, x, y)]
}

// Squarei gets the field value at the specified index.
func (g *Grid) Squarei(i int) Square {
	if !indexExists(g, i) {
		panic("grid.Squarei: index out of range.")
	}
	return g.Squares[i]
}

// SetSquare sets the squares value at coordinates x, y.
func (g *Grid) SetSquare(x int, y int, val Square) *Grid {
	if !coordsExist(g, x, y) {
		panic("grid.SetSquare: coords out of range.")
	}
	g.Squares[coordsToIndex(g, x, y)] = val
	return g
}

// SetSquarei sets the squares value at the specified index.
func (g *Grid) SetSquarei(i int, val Square) *Grid {
	if !indexExists(g, i) {
		panic("grid.SetSquarei: index out of range.")
	}
	g.Squares[i] = val
	return g
}

// SetSquareiAndValidate sets a square to the specified value and validates the grid partially (only the changed square and its neighbors).
func (g *Grid) SetSquareiAndValidate(i int, val Square) bool {
	g.SetSquarei(i, val)
	ixs := append(g.NeighborIndicesi(i, false), i)
	return ValidatePartial(g, ixs)
}

// HasSquare returns "true" if at least one square with the specified value exists.
func (g *Grid) HasSquare(val Square) bool {
	for _, v := range g.Squares {
		if v == val {
			return true
		}
	}
	return false
}

// Fill fills all squares of a grid with the specified value.
//
// TODO: move to Builder type
func (g *Grid) Fill(s Square) *Grid {
	for i := range g.Squares {
		g.Squares[i] = s
	}
	return g
}

// FillUndefined fills all undefined squares of a grid with the specified value.
//
// TODO: move to Builder type
func (g *Grid) FillUndefined(s Square) *Grid {
	for i := range g.Squares {
		if g.Squares[i] == SquareUndefined {
			g.Squares[i] = s
		}
	}
	return g
}

// Neighbors gets the neighbours field values.
func (g *Grid) Neighbors(x int, y int) []Square {
	return []Square{
		g.Square(x-1, y-1),
		g.Square(x, y-1),
		g.Square(x+1, y-1),
		g.Square(x-1, y),
		g.Square(x+1, y),
		g.Square(x-1, y+1),
		g.Square(x, y+1),
		g.Square(x+1, y+1),
	}
}

// NeighborIndicesi gets the indices of all neighbor squares.
func (g *Grid) NeighborIndicesi(i int, adjacentOnly bool) []int {
	res := make([]int, 0, maxNumNeighbors)
	x, y := g.Coords(i)
	for _, n := range neighbours {
		if (!adjacentOnly || n.adjacent) && coordsExist(g, x+n.x, y+n.y) {
			res = append(res, coordsToIndex(g, x+n.x, y+n.y))
		}
	}
	return res
}

// GetAdjacentNeighbors gets the adjacent neighbours field values.
func (g *Grid) GetAdjacentNeighbors(x int, y int) []Square {
	return []Square{
		g.Square(x, y-1),
		g.Square(x-1, y),
		g.Square(x+1, y),
		g.Square(x, y+1),
	}
}

// NeighborCount counts the neighboured squares (all 8 or the 4 adjacent) that match the given type.
// If the 'includeUndefined' flag is set, also undefined squares are counted.
func (g *Grid) NeighborCount(i int, square Square, adjacentOnly bool, includeUndefined bool) (count int) {
	for _, ni := range g.NeighborIndicesi(i, adjacentOnly) {
		v := g.Squarei(ni)
		if v == square || (includeUndefined && v == SquareUndefined) {
			count++
		}
	}
	return
}

// CountNeighbors counts the neighboured squares that match the given type.
//
// Deprecated: use NeighborCount
func (g *Grid) CountNeighbors(i int, square Square) int {
	return g.NeighborCount(i, square, false, false)
}

// CountAdjacentNeighbours counts the adacent neighboured squares that match the given type.
//
// Deprecated: use NeighborCount
func (g *Grid) CountAdjacentNeighbours(i int, square Square) int {
	return g.NeighborCount(i, square, true, false)
}

// Clone creates an exact deep copy of a grid.
func (g *Grid) Clone() *Grid {
	clone := *g
	clone.Squares = make([]Square, len(g.Squares))
	copy(clone.Squares, g.Squares)
	return &clone
}

// SetDragon sets a dragon to a specific square and computes where fire must be.
//
// TODO: move this into a Builder type
func (g *Grid) SetDragon(i int) *Grid {
	g.SetSquarei(i, SquareDragon)
	for _, ni := range g.NeighborIndicesi(i, false) {
		if g.CountNeighbors(ni, SquareDragon) > 1 {
			g.SetSquarei(ni, SquareFire)
		}
	}
	return g
}

// Interestingness returns the interestingness of a grid.
// This is useful if we want to sort possible solved grids.
func (g *Grid) Interestingness() (interestingness int) {
	for _, val := range g.Squares {
		interestingness += squareInterestingness[val]
	}
	return
}

// Size returns a grids number of squares.
func (g *Grid) Size() int {
	return g.Width * g.Height
}

// CountSquares counts all squares with a specified value.
func (g *Grid) CountSquares(s Square) (count int) {
	for _, v := range g.Squares {
		if v == s {
			count++
		}
	}
	return
}

// IsUndefined return true if all squares are undefined
func (g *Grid) IsUndefined() bool {
	return g.Size() == g.CountSquares(SquareUndefined)
}

// String returns the string representation of a puzzle.
func (g *Grid) String() string {
	sb := strings.Builder{}
	sb.WriteString("   ┌")
	sb.WriteString(strings.Repeat("─", 2*g.Width+1))
	sb.WriteString("┐\n")
	for i, val := range g.Squares {
		if i%g.Width == 0 {
			sb.WriteString("   │ ")
		}
		sb.WriteRune(getSymbol(val))
		sb.WriteRune(' ')
		if i%g.Width == g.Width-1 {
			sb.WriteString("│")

			if i/g.Width == 0 {
				sb.WriteString(fmt.Sprintf(" Size: %vx%v", g.Width, g.Height))
			} else if i/g.Width == 1 {
				sb.WriteString(" Code: ")
				for i, val := range g.Squares {
					sb.WriteRune(getSymbolForCode(val))
					if i%g.Width == g.Width-1 && i < g.Width*g.Height-1 {
						sb.WriteString(",")
					}
				}
			}

			sb.WriteString("\n")
		}
	}
	sb.WriteString("   └")
	sb.WriteString(strings.Repeat("─", 2*g.Width+1))
	sb.WriteString("┘")
	return sb.String()
}

func coordsToIndex(g *Grid, x int, y int) int {
	return g.Width*y + x
}

func coordsExist(g *Grid, x int, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func indexExists(g *Grid, i int) bool {
	return i >= 0 && i < g.Width*g.Height
}

func getSymbol(val Square) rune {
	return squareSymbols[val]
}

func getSymbolForCode(val Square) rune {
	return squareSymbolsForCode[val]
}
