package model

import "testing"

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
		g.Neighborsi(7, false)
	}
}

func BenchmarkNeighborsiAdjacent(b *testing.B) {
	g := Parse("__xdx_,xf____,_fd_xd,_f____,xfx_dx,x_d_x_")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.Neighborsi(7, true)
	}
}
