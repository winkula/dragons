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

// ParseWorld parses a string that represents a dragons world.
func ParseWorld(s string) *World {
	matched, _ := regexp.MatchString(`^[0-9]+,[0-9]+$`, s)
	if matched {
		parts := strings.Split(s, delimiter)
		width, _ := strconv.ParseInt(parts[0], 10, 32)
		height, _ := strconv.ParseInt(parts[1], 10, 32)
		return NewWorld(int(width), int(height))
	}

	parts := strings.Split(s, delimiter)
	height := len(parts)
	if height == 0 {
		panic("Parse error: World has no height, invlid.")
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

	world := NewWorld(width, height)
	world.Squares = squares
	return world
}

func getSquare(square rune) Square {
	switch unicode.ToLower(square) {
	case '_', '.':
		return SquareUndefined
	case 'x':
		return SquareEmpty
	case 'f':
		return SquareFire
	case 'd':
		return SquareDragon
	}
	panic(fmt.Sprintf("Parse error: '%c' is not a valid value.", square))
}
