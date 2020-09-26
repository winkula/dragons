package model

import "testing"

func TestValidate(t *testing.T) {
	tables := []struct {
		grid  *Grid
		valid bool
		rule  string
	}{
		// The territory rule is fulfilled.
		{New(3, 3), true, "territory"},
		{New(3, 3).
			SetSquare(0, 0, SquareDragon), true, "territory"},
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(1, 1, SquareFire).
			SetSquare(2, 2, SquareDragon), true, "territory"},
		// The territory rule is violated.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(1, 1, SquareDragon), false, "territory"},
		{New(3, 3).
			SetSquare(1, 0, SquareDragon).
			SetSquare(2, 0, SquareDragon), false, "territory"},
		// The survive rule is fulfilled.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon), true, "survive"},
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(1, 1, SquareFire), true, "survive"},
		// The survive rule is violated.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(1, 0, SquareFire).
			SetSquare(0, 1, SquareFire), false, "survive"},
		{New(3, 3).
			SetSquare(1, 1, SquareDragon).
			SetSquare(0, 0, SquareFire).
			SetSquare(1, 0, SquareFire).
			SetSquare(2, 0, SquareFire).
			SetSquare(0, 1, SquareFire).
			SetSquare(2, 1, SquareFire), false, "survive"},
		// The fight rule is fulfilled.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(2, 2, SquareDragon).
			SetSquare(1, 1, SquareFire), true, "fight"},
		// The fight rule is violated.
		{New(3, 3).
			SetSquare(0, 0, SquareDragon).
			SetSquare(2, 1, SquareDragon).
			SetSquare(1, 0, SquareEmpty).
			SetSquare(1, 1, SquareFire), false, "fight"},
		{New(1, 1).
			SetSquare(0, 0, SquareFire), false, "fight"},
	}

	for _, table := range tables {
		valid := Validate(table.grid)
		if valid != table.valid {
			t.Errorf("Validate was incorrect, got: %t, want: %t (%s rule). Grid: \n%s",
				valid,
				table.valid,
				table.rule,
				table.grid)
		}

		valid = ValidateFast(table.grid)
		if valid != table.valid {
			t.Errorf("ValidateFast was incorrect, got: %t, want: %t (%s rule). Grid: \n%s",
				valid,
				table.valid,
				table.rule,
				table.grid)
		}
	}
}

// History
// - 1904 ns/op
// - 1591 ns/op (after switching from int to uint8 for square values)
// - 1392 ns/op (after reducing multiple calls to CountNeighbors)
func BenchmarkValidate(b *testing.B) {
	gs := []*Grid{
		Parse("_xf_,____,____,_d__"),
		Parse("_f_f_,df_f_,_fxf_,x___x,_d_d_"),
		Parse("_____,d__f_,_____,_____,__f__"),
		Parse("__xdx_,xf____,_fd_xd,_f____,xfx_dx,x_d_x_"),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gs[i%len(gs)]
		Validate(w)
	}
}

func BenchmarkValidateFast(b *testing.B) {
	gs := []*Grid{
		Parse("_xf_,____,____,_d__"),
		Parse("_f_f_,df_f_,_fxf_,x___x,_d_d_"),
		Parse("_____,d__f_,_____,_____,__f__"),
		Parse("__xdx_,xf____,_fd_xd,_f____,xfx_dx,x_d_x_"),
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := gs[i%len(gs)]
		ValidateFast(w)
	}
}
