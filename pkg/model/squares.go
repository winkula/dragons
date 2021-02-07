package model

// Square represents possible cell values.
type Square uint8

const (
	// SquareUndefined represents undefined squares.
	SquareUndefined Square = iota
	// SquareOut represents squares that are outside of the grid.
	SquareOut
	// SquareAir represents squares with air.
	SquareAir
	// SquareFire represents squares with fire.
	SquareFire
	// SquareDragon represents squares with dragons.
	SquareDragon
)

const numSquareValues = 5

// AllFields are a list of all possible field values.
var AllFields = []Square{
	SquareUndefined,
	SquareOut,
	SquareAir,
	SquareFire,
	SquareDragon,
}

var squareAttributes = map[Square]struct {
	symbol      rune
	code        rune
	density     int
	puzzleValue int
}{
	SquareUndefined: {' ', '_', 0, 100},
	SquareOut:       {'#', '#', 0, 0},
	SquareAir:       {'-', 'x', 0, -1},
	SquareFire:      {'Δ' /* 🔥 */, 'f', 1, -5},
	SquareDragon:    {'▲' /* 🐲*/, 'd', 1, -50},
}

func (val Square) Symbol() rune {
	return squareAttributes[val].symbol
}

func (val Square) Code() rune {
	return squareAttributes[val].code
}

func (val Square) Density() int {
	return squareAttributes[val].density
}

func (val Square) PuzzleValue() int {
	return squareAttributes[val].puzzleValue
}
