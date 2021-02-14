package model

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	delimiter = ","
)

// Parse parses a string that represents a grid.
func Parse(s string) *Grid {
	matched, _ := regexp.MatchString(`^[0-9]+,[0-9]+$`, s)
	if matched {
		parts := strings.Split(s, delimiter)
		width, _ := strconv.ParseInt(parts[0], 10, 32)
		height, _ := strconv.ParseInt(parts[1], 10, 32)
		return New(int(width), int(height))
	}

	parts := strings.Split(s, delimiter)
	height := len(parts)
	if height == 0 {
		panic("Parse error: grid has no height, invalid.")
	}

	width := len(parts[0])
	squares := make([]Square, 0, width*height)
	for _, row := range parts {
		if len(row) != width {
			panic("Parse error: Not all rows have the same width")
		}
		for _, c := range row {
			squares = append(squares, getSquare(c))
		}
	}

	g := New(width, height)
	g.Squares = squares
	return g
}

// TODO: use squareAttributes here
func getSquare(square rune) Square {
	switch unicode.ToLower(square) {
	case '_', '.':
		return SquareUndefined
	case 'x':
		return SquareAir
	case 'f':
		return SquareFire
	case 'd':
		return SquareDragon
	case 'n':
		return SquareNoDragon
	}
	panic(fmt.Sprintf("Parse error: '%c' is not a valid value.", square))
}
