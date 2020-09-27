package model

import (
	"testing"
)

func TestEnumerate(t *testing.T) {
	tables := []struct {
		grid  *Grid
		count int
	}{
		{New(1, 1), 1},
		{New(1, 2), 1},
		{New(2, 1), 1},
		{New(2, 2), 5},
		{New(2, 3), 7},
		{New(3, 2), 7},
		{New(3, 3), 14},

		{Parse("____,_ff_,_xd_"), 2},
		{Parse("____,_ffx,_xd_"), 1},
	}

	for _, table := range tables {
		sucs := Enumerate(table.grid)
		length := len(sucs)
		if length != table.count {
			t.Errorf("Enumerate was incorrect, got: %d, want: %d. Grid: \n%s",
				length,
				table.count,
				table.grid)
		}
	}
}

func TestEnumerateSquare(t *testing.T) {
	tables := []struct {
		grid  *Grid
		index int
		count int
	}{
		{Parse("x_,xd"), 1, 1},
		{Parse("xx,x_"), 3, 2},
		{Parse("d_,__"), 0, 1}, // only one possible: enumerating an already defined square
	}

	for _, table := range tables {
		sucs := EnumerateSquare(table.grid, table.index)
		length := len(sucs)
		if length != table.count {
			t.Errorf("Enumerate was incorrect, got: %d, want: %d. Grid: \n%s",
				length,
				table.count,
				table.grid)
		}
	}
}

func TestIsDistinct(t *testing.T) {
	tables := []struct {
		grid     *Grid
		expected bool
	}{
		{Parse("___,d__,___"), false},
		{Parse("___,d_d,___"), true},
		{Parse("___,dd_,___"), false}, // invalid grid
	}

	for _, table := range tables {
		result := IsDistinct(table.grid)

		if result != table.expected {
			t.Errorf("IsDistinct was incorrect, got: %t, want: %t. Grid: \n%s",
				result,
				table.expected,
				table.grid)
		}
	}
}

// history:
// - 9077091 ns/op
// - 4670463 ns/op
func BenchmarkEnumerate(b *testing.B) {
	g := New(4, 4)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Enumerate(g)
	}
}
