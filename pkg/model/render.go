package model

import (
	"fmt"
	"strings"
)

// Render returns the string representation of a grid and shows some extra information.
func Render(g *Grid, k *knowledge, activeSquare int) string {
	sb := strings.Builder{}
	sb.WriteString("   ┌")
	sb.WriteString(strings.Repeat("─", 2*g.Width+1))
	sb.WriteString("┐\n")
	for i, val := range g.Squares {
		if i%g.Width == 0 {
			sb.WriteString("   │")
			if i == activeSquare {
				sb.WriteRune('[')
			} else {
				sb.WriteRune(' ')
			}
		}
		sb.WriteRune(getSymbol(val))

		if i == activeSquare {
			sb.WriteRune(']')
		} else if i == activeSquare-1 && i%g.Width != g.Width-1 {
			sb.WriteRune('[')
		} else {
			sb.WriteRune(' ')
		}

		if i%g.Width == g.Width-1 {
			sb.WriteString("│")

			// additional information on the right side
			if i/g.Width == 0 {
				sb.WriteString(fmt.Sprintf(" Size: %vx%v", g.Width, g.Height))
			} else if i/g.Width == 1 {
				sb.WriteString(" Code: ")
				for i, val := range g.Squares {
					sb.WriteRune(getSymbolForCode(val))
					if i%g.Width == g.Width-1 && i < g.Width*g.Height-1 {
						sb.WriteString(",")
					}
				}
			}

			sb.WriteString("\n")
		}
	}
	sb.WriteString("   └")
	sb.WriteString(strings.Repeat("─", 2*g.Width+1))
	sb.WriteString("┘")
	return sb.String()
}
