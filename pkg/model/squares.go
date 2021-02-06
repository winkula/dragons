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

var squareSymbols = map[Square]rune{
	SquareUndefined: ' ',
	SquareOut:       '#',
	SquareAir:       '-',
	SquareFire:      'Δ', // 🔥
	SquareDragon:    '▲', // 🐲
}

var squareSymbolsForCode = map[Square]rune{
	SquareUndefined: '_',
	SquareOut:       '#',
	SquareAir:       'x',
	SquareFire:      'f',
	SquareDragon:    'd',
}

var squareDensity = map[Square]int{
	SquareUndefined: 0,
	SquareOut:       0,
	SquareAir:       0,
	SquareFire:      1,
	SquareDragon:    1,
}
