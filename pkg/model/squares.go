package model

// Square represents possible cell values.
type Square uint8

const (
	// SquareUndefined represents undefined squares.
	SquareUndefined Square = iota
	// SquareAir represents squares with air.
	SquareAir
	// SquareFire represents squares with fire.
	SquareFire
	// SquareDragon represents squares with dragons.
	SquareDragon
	// SquareOut represents squares that are outside of the grid.
	SquareOut
)

const numSquareValues = 5

// AllFields are a list of all possible field values.
var AllFields = []Square{
	SquareUndefined,
	SquareAir,
	SquareFire,
	SquareDragon,
	SquareOut,
}

var squareAttributes = map[Square]struct {
	symbol      rune
	code        rune
	density     int
	puzzleValue int
}{
	SquareUndefined: {' ', '_', 0, 10},
	SquareAir:       {'-', 'x', 0, 9},
	SquareFire:      {'Δ' /* 🔥 */, 'f', 1, 3},
	SquareDragon:    {'▲' /* 🐲*/, 'd', 1, 1},
	SquareOut:       {'#', '#', 0, 0},
}

// Symbol is the squares representation in the console output and logs.
func (val Square) Symbol() rune {
	return squareAttributes[val].symbol
}

// Code is the squares representation in when serializing grids.
func (val Square) Code() rune {
	return squareAttributes[val].code
}

// Density is the squares density value.
func (val Square) Density() int {
	return squareAttributes[val].density
}

// PuzzleValue is the squares puzzle value.
func (val Square) PuzzleValue() int {
	return squareAttributes[val].puzzleValue
}
