package model

import "testing"

func TestSolve(t *testing.T) {
	tables := []*Grid{
		Parse("dx,_x"),
		Parse("dx,__"),
		Parse("_f_,_f_,_f_"),
		Parse("_f_,___,___"),
		Parse("d__,___,__d"),
	}

	for _, table := range tables {
		solved := Solve(table)
		if solved == nil || !isSolved(solved) {
			t.Errorf("TestSolve was incorrect, grid is not solved. Grid: \n%s",
				table)
		}
	}
}

func BenchmarkSolve(b *testing.B) {
	g := Parse("_f_,_f_,_f_")
	for i := 0; i < b.N; i++ {
		solved := Solve(g)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
		}
	}
}

func TestSolveDk(t *testing.T) {
	tables := []*Grid{
		Parse("dx,_x"),
		Parse("dx,__"),
		Parse("_f_,_f_,_f_"),
		//Parse("_f_,___,___"), // not solvable
		Parse("d__,___,__d"),
	}

	for _, table := range tables {
		solved := SolveDk(table)
		if solved == nil || !isSolved(solved) {
			t.Errorf("TestSolve was incorrect, grid is not solved. Grid: \n%s",
				table)
		}
	}
}

func BenchmarkSolveDk(b *testing.B) {
	g := Parse("_f_,_f_,_f_")
	for i := 0; i < b.N; i++ {
		solved := SolveDk(g)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
		}
	}
}

func TestSolveBf(t *testing.T) {
	tables := []*Grid{
		Parse("dx,_x"),
		Parse("dx,__"),
		Parse("_f_,_f_,_f_"),
		Parse("_f_,___,___"),
		Parse("d__,___,__d"),
	}

	for _, table := range tables {
		solved := SolveBf(table)
		if solved == nil || !isSolved(solved) {
			t.Errorf("TestSolve was incorrect, grid is not solved. Grid: \n%s",
				table)
		}
	}
}

func BenchmarkSolveBf(b *testing.B) {
	g := Parse("_f_,_f_,_f_")
	for i := 0; i < b.N; i++ {
		solved := SolveBf(g)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
		}
	}
}
