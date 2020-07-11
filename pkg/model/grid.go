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

func coordsToIndex(g *Grid, x int, y int) int {
	return g.Width*y + x
}

func coordsExist(g *Grid, x int, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func indexExists(g *Grid, i int) bool {
	return i >= 0 && i < g.Width*g.Height
}

// Coords gets the coordinates of a quare given the squares index.
func (g *Grid) Coords(i int) (x int, y int) {
	return i % g.Width, i / g.Width
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
		panic(fmt.Sprintf("Error: invalid index (index = %v, width = %v, height = %v)", i, g.Width, g.Height))
	}
	return g.Squares[i]
}

// SetSquare sets the squares value at coordinates x, y.
func (g *Grid) SetSquare(x int, y int, val Square) *Grid {
	if !coordsExist(g, x, y) {
		panic(fmt.Sprintf("Error: invalid coordinates (x = %v, y = %v, width = %v, height = %v)", x, y, g.Width, g.Height))
	}
	g.Squares[coordsToIndex(g, x, y)] = val
	return g
}

// SetSquarei sets the squares value at the specified index.
func (g *Grid) SetSquarei(i int, val Square) *Grid {
	if !indexExists(g, i) {
		panic(fmt.Sprintf("Error: invalid index (index = %v, width = %v, height = %v)", i, g.Width, g.Height))
	}
	g.Squares[i] = val
	return g
}

// FillSquares fills all squares.
func (g *Grid) FillSquares(val Square) *Grid {
	for i := range g.Squares {
		g.Squares[i] = val
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

// Neighborsi gets the neighbours indexes.
func (g *Grid) Neighborsi(i int, adjacent bool) []int {
	res := make([]int, 0, maxNumNeighbors)
	x, y := g.Coords(i)
	for _, n := range neighbours {
		if (!adjacent || n.adjacent) && coordsExist(g, x+n.x, y+n.y) {
			res = append(res, coordsToIndex(g, x+n.x, y+n.y))
		}
	}
	return res
}

// GetAdjacentNeighbours gets the adjacent neighbours field values.
func (g *Grid) GetAdjacentNeighbours(x int, y int) []Square {
	return []Square{
		g.Square(x, y-1),
		g.Square(x-1, y),
		g.Square(x+1, y),
		g.Square(x, y+1),
	}
}

// CountNeighbours counts the neighboured squares that match the given type.
func (g *Grid) CountNeighbours(i int, square Square) int {
	x, y := g.Coords(i)
	ns := g.Neighbors(x, y)
	count := 0
	for _, v := range ns {
		if v == square {
			count++
		}
	}
	return count
}

// CountAdjacentNeighbours counts the adacent neighboured squares that match the given type.
func (g *Grid) CountAdjacentNeighbours(i int, square Square) int {
	x, y := g.Coords(i)
	ns := g.GetAdjacentNeighbours(x, y)
	count := 0
	for _, v := range ns {
		if v == square {
			count++
		}
	}
	return count
}

// Clone creates an exact deep copy of a grid.
func (g *Grid) Clone() *Grid {
	clone := *g
	clone.Squares = make([]Square, len(g.Squares))
	copy(clone.Squares, g.Squares)
	return &clone
}

// SetDragon sets a dragon to a specific square and computes where fire must be.
func (g *Grid) SetDragon(i int) {
	x, y := g.Coords(i)
	g.SetSquare(x, y, SquareDragon)
	ns := g.Neighbors(x, y)
	for i := range ns {
		numDragons := g.CountNeighbours(i, SquareDragon)
		if numDragons > 1 {
			g.SetSquarei(i, SquareFire)
		}
	}
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

// CountSquares fills all squares.
func (g *Grid) CountSquares(needle Square) int {
	return countSquares(g.Squares, needle)
}

// String returns the string representation of a puzzle.
func (g *Grid) String() string {
	sb := strings.Builder{}
	sb.WriteString("   ┌")
	if g.Width > 1 {
		pad := (2*g.Width + 3 - 7) / 2
		sb.WriteString(strings.Repeat("─", pad))
		sb.WriteString(fmt.Sprintf("╴%vx%v╶", g.Width, g.Height))
		sb.WriteString(strings.Repeat("─", pad))
	} else {
		sb.WriteString(strings.Repeat("─", 2*g.Width+1))
	}
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
