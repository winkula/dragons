package model

// Square represents possible cell values.
type Square int

const (
	// SquareUndefined represent undefined squares.
	SquareUndefined Square = iota
	// SquareOut represent undefined squares.
	SquareOut
	// SquareEmpty represent empty squares.
	SquareEmpty
	// SquareFire represent squares with fire.
	SquareFire
	// SquareDragon represent squares with dragons.
	SquareDragon
)

var squareSymbols = map[Square]rune{
	SquareUndefined: '.',
	SquareOut:       ' ',
	SquareEmpty:     'x',
	SquareFire:      'F',
	SquareDragon:    'D',
}
