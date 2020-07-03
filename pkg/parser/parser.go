package parser

import (
	"github.com/winkula/dragons/pkg/model"
)

// ParseWorld parses a string that represents a dragons world.
func ParseWorld(s string) *model.World {
	width := 0
	height := 0
	squares := make([]model.Square, 0)
	for _, c := range s {
		if isDelim(c) {
			height++
		} else {
			square := getSquare(c)
			squares = append(squares, square)
			if height == 0 {
				width++
			}
		}
	}
	world := model.NewWorld(width, height+1)
	world.Squares = squares
	return world
}

func isDelim(char rune) bool {
	return char == ','
}

func getSquare(square rune) model.Square {
	switch square {
	case 'x':
	case 'X':
		return model.SquareEmpty
	case 'f':
	case 'F':
		return model.SquareFire
	case 'd':
	case 'D':
		return model.SquareDragon
	}
	return model.SquareUndefined
}
