package model

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
	if !g.coordsExist(x, y) {
		return -1, false // index does not exist
	}
	i := g.coordsToIndex(x, y)
	return i, true
}

// Square gets the field value at coordinates x, y.
func (g *Grid) Square(x int, y int) Square {
	if !g.coordsExist(x, y) {
		return SquareOut
	}
	return g.Squares[g.coordsToIndex(x, y)]
}

// Squarei gets the field value at the specified index.
func (g *Grid) Squarei(i int) Square {
	if !g.indexExists(i) {
		panic("grid.Squarei: index out of range.")
	}
	return g.Squares[i]
}

// SetSquare sets the squares value at coordinates x, y.
func (g *Grid) SetSquare(x int, y int, val Square) *Grid {
	if !g.coordsExist(x, y) {
		panic("grid.SetSquare: coords out of range.")
	}
	g.Squares[g.coordsToIndex(x, y)] = val
	return g
}

// SetSquarei sets the squares value at the specified index.
func (g *Grid) SetSquarei(i int, val Square) *Grid {
	if !g.indexExists(i) {
		panic("grid.SetSquarei: index out of range.")
	}
	g.Squares[i] = val
	return g
}

// SetSquareiAndValidate sets a square to the specified value and validates the grid partially (only the changed square and its neighbors).
//
// TODO: move to builde type
func (g *Grid) SetSquareiAndValidate(i int, val Square) bool {
	g.SetSquarei(i, val)
	return ValidateIncr(g, i, 1)
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

// NeighborIndicesi gets the indices of all neighbor squares.
func (g *Grid) NeighborIndicesi(i int, adjacentOnly bool) []int {
	res := make([]int, 0, maxNumNeighbors)
	x, y := g.Coords(i)
	for _, n := range neighbours {
		if (!adjacentOnly || n.adjacent) && g.coordsExist(x+n.x, y+n.y) {
			res = append(res, g.coordsToIndex(x+n.x, y+n.y))
		}
	}
	return res
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

// NeighborCount8 counts the 8 neighboured squares that match the given type.
func (g *Grid) NeighborCount8(i int, square Square) int {
	return g.NeighborCount(i, square, false, false)
}

// NeighborCount4 counts the adacent neighboured squares that match the given type.
func (g *Grid) NeighborCount4(i int, square Square) int {
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
		if g.NeighborCount8(ni, SquareDragon) > 1 {
			g.SetSquarei(ni, SquareFire)
		}
	}
	return g
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
	return Render(g, -1)
}

func (g *Grid) coordsToIndex(x int, y int) int {
	return g.Width*y + x
}

func (g *Grid) coordsExist(x int, y int) bool {
	return x >= 0 && x < g.Width && y >= 0 && y < g.Height
}

func (g *Grid) indexExists(i int) bool {
	return i >= 0 && i < g.Width*g.Height
}
