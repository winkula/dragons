package game

import (
	"fmt"
	"time"

	"github.com/winkula/dragons/pkg/model"
)

// Game represents a dragons puzzle interactive game.
type Game struct {
	Puzzle   *model.Grid
	Solution *model.Grid
	State    *model.Grid
	CursorX  int
	CursorY  int
}

// NewGame creates a new game.
func NewGame(width int, height int, duration time.Duration) *Game {
	solution := model.Generate(width, height, time.Duration(duration/2))
	puzzle := model.Obfuscate(solution, model.DifficultyEasy, time.Duration(duration/2))

	return &Game{
		Puzzle:   puzzle,
		Solution: solution,
		State:    puzzle.Clone(),
		CursorX:  0,
		CursorY:  0,
	}
}

// NewGameFromPuzzle creates a new game from a given puzzle.
func NewGameFromPuzzle(g *model.Grid) *Game {
	puzzle := g
	solution := model.Solve(puzzle)

	return &Game{
		Puzzle:   puzzle,
		Solution: solution,
		State:    puzzle.Clone(),
		CursorX:  0,
		CursorY:  0,
	}
}

func (g *Game) Left() {
	if g.CursorX > 0 {
		g.CursorX--
	}
}

func (g *Game) Right() {
	if g.CursorX < g.State.Width-1 {
		g.CursorX++
	}
}

func (g *Game) Up() {
	if g.CursorY > 0 {
		g.CursorY--
	}
}

func (g *Game) Down() {
	if g.CursorY < g.State.Height-1 {
		g.CursorY++
	}
}

func (g *Game) Set(square model.Square) {
	index := g.State.Height*g.CursorY + g.CursorX

	if g.Puzzle.Squares[index] == model.SquareUndefined {
		g.State.Squares[index] = square
	}
}

func (g *Game) IsValid() bool {
	for i := 0; i < g.Puzzle.Size(); i++ {
		if g.State.Squares[i] != model.SquareUndefined && g.State.Squares[i] != g.Solution.Squares[i] {
			return false
		}
	}
	return true
}

func (g *Game) IsSolved() bool {
	for i := 0; i < g.Puzzle.Size(); i++ {
		if g.State.Squares[i] != g.Solution.Squares[i] {
			return false
		}
	}
	return true
}

func (g *Game) String() string {
	if g.IsSolved() {
		return " SOLVED"
	}
	if g.IsValid() {
		return " VALID"
	}
	return fmt.Sprintf(" HAS ERRORS")
}
