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
	// SquareNoDragon represents squares that can not be dragons.
	SquareNoDragon
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
	symbol rune
	code   rune
}{
	SquareUndefined: {' ', '_'},
	SquareAir:       {'-', 'x'},
	SquareFire:      {'Δ' /* 🔥 */, 'f'},
	SquareDragon:    {'▲' /* 🐲*/, 'd'},
	SquareOut:       {'#', '#'},
}

// Symbol is the squares representation in the console output and logs.
func (val Square) Symbol() rune {
	return squareAttributes[val].symbol
}

// Code is the squares representation in when serializing grids.
func (val Square) Code() rune {
	return squareAttributes[val].code
}
