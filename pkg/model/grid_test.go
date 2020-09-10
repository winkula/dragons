package model

import "testing"

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
