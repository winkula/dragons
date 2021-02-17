package model

import (
	"math/big"
	"strings"
	"testing"
)

func TestNewGrid(t *testing.T) {
	tables := []struct {
		width  int
		height int
	}{
		{2, 4},
		{3, 3},
	}

	for _, table := range tables {
		grid := New(table.width, table.height)

		if grid.Width != table.width {
			t.Errorf("TestNewGrid was incorrect, got: %d, want: %d. Grid: \n%s",
				grid.Width, table.width, grid)
		}
		if grid.Height != table.height {
			t.Errorf("TestNewGrid was incorrect, got: %d, want: %d. Grid: \n%s",
				grid.Height, table.height, grid)
		}
	}
}

func TestSquare(t *testing.T) {
	tables := []struct {
		grid     *Grid
		x        int
		y        int
		expected Square
	}{
		{Parse("dx,xx"), 0, 0, SquareDragon},
		{Parse("dx,xx"), 1, 1, SquareAir},
		{Parse("dx,xx"), 5, 5, SquareOut},
	}

	for _, table := range tables {
		result := table.grid.Square(table.x, table.y)

		if result != table.expected {
			t.Errorf("TestSquare was incorrect, got: %d, want: %d. Grid: \n%s",
				result, table.expected, table.grid)
		}
	}
}

func TestString(t *testing.T) {
	tables := []struct {
		grid              *Grid
		expectedSubstring string
	}{
		{Parse("dx,xx"), "Size: 2x2"},
		{Parse("dx,xx"), "Code: dx,xx"},
		{Parse("dx,xx"), "│ ▲ - │"},
		{Parse("xfx,dfd,xfx"), "│ ▲ Δ ▲ │"},
	}

	for _, table := range tables {
		result := table.grid.String()

		if !strings.Contains(result, table.expectedSubstring) {
			t.Errorf("TestString was incorrect, got: %s, want: %s. Grid: \n%s",
				result, table.expectedSubstring, table.grid)
		}
	}
}

func TestIndex(t *testing.T) {
	tables := []struct {
		grid  *Grid
		x     int
		y     int
		index int
		ok    bool
	}{
		{New(3, 3), 0, 0, 0, true},
		{New(3, 3), 1, 1, 4, true},
		{New(3, 3), 0, 3, -1, false},
		{New(3, 3), 3, 0, -1, false},
	}

	for _, table := range tables {
		i, ok := table.grid.Index(table.x, table.y)

		if i != table.index {
			t.Errorf("Index was incorrect, got: %v, want: %v. Grid: \n%s",
				i, table.index, table.grid)
		}
		if ok != table.ok {
			t.Errorf("Index was incorrect, got: %v, want: %v. Grid: \n%s",
				i, table.ok, table.grid)
		}
	}
}

func TestSolutionRating(t *testing.T) {
	tables := []struct {
		grid   *Grid
		rating float64
	}{
		{Parse("__,__"), 0.0},
		{Parse("dx,_f"), 0.25},
	}
	for _, table := range tables {
		rating := table.grid.SolutionRating()

		if table.rating != rating {
			t.Errorf("SolutionRating was incorrect, got: %v. Expected %v", rating, table.rating)
		}
	}
}

func TestID(t *testing.T) {
	tables := []struct {
		grid *Grid
		code int64
	}{
		{Parse("__,__"), 0b_00010_00010},
		{Parse("x"), 0b_01_00001_00001},
		{Parse("f"), 0b_10_00001_00001},
		{Parse("d"), 0b_11_00001_00001},
		{Parse("d__"), 0b_11_00001_00011},
		{Parse("__d"), 0b_11_00_00_00001_00011},
		{Parse("dx,_f"), 0b_10_00_01_11_00010_00010},
	}
	for _, table := range tables {
		code := table.grid.ID()

		if big.NewInt(table.code).Cmp(code) != 0 {
			t.Errorf("Code was incorrect, got: %v. Expected %v", code, table.code)
		}
	}
}

func BenchmarkNeighborsi(b *testing.B) {
	g := Parse("__xdx_,xf____,_fd_xd,_f____,xfx_dx,x_d_x_")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.NeighborIndicesi(7, false)
	}
}

func BenchmarkNeighborsiAdjacent(b *testing.B) {
	g := Parse("__xdx_,xf____,_fd_xd,_f____,xfx_dx,x_d_x_")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.NeighborIndicesi(7, true)
	}
}
