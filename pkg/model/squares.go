package model

// Square represents possible cell values.
type Square uint8

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

const numSquareValues = 5

var AllFields = []Square{
	SquareUndefined,
	SquareOut,
	SquareEmpty,
	SquareFire,
	SquareDragon,
}

var squareSymbols = map[Square]rune{
	SquareUndefined: ' ',
	SquareOut:       '#',
	SquareEmpty:     '-',
	SquareFire:      'Δ', // 🔥
	SquareDragon:    '▲', // 🐲
}

var squareSymbolsForCode = map[Square]rune{
	SquareUndefined: '_',
	SquareOut:       '#',
	SquareEmpty:     'x',
	SquareFire:      'f',
	SquareDragon:    'd',
}

var squareDensity = map[Square]int{
	SquareUndefined: 0,
	SquareOut:       0,
	SquareEmpty:     0,
	SquareFire:      1,
	SquareDragon:    1,
}
