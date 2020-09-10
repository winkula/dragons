package model

import (
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
		{Parse("dx,xx"), 1, 1, SquareEmpty},
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
		{Parse("dx,xx"), "│ D - │"},
	}

	for _, table := range tables {
		result := table.grid.String()

		if !strings.Contains(result, table.expectedSubstring) {
			t.Errorf("TestString was incorrect, got: %s, want: %s. Grid: \n%s",
				result, table.expectedSubstring, table.grid)
		}
	}
}

func BenchmarkNeighbors(b *testing.B) {
	g := Parse("__xdx_,xf____,_fd_xd,_f____,xfx_dx,x_d_x_")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Neighbors(1, 1)
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
