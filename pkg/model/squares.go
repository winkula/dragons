package model

// Square represents possible cell values.
type Square int

const (
	// SquareUndefined represent undefined squares.
	SquareUndefined Square = iota
	// SquareOut represent undefined squares.
	SquareOut
	// SquareFire represent squares with fire.
	SquareFire
	// SquareDragon represent squares with dragons.
	SquareDragon
)
