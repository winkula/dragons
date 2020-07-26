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
		solved, _ := Solve(table)
		if solved == nil || !isSolved(solved) {
			t.Errorf("TestSolve was incorrect, grid is not solved. Grid: \n%s",
				table)
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
		solved, _ := SolveDk(table)
		if solved == nil || !isSolved(solved) {
			t.Errorf("TestSolve was incorrect, grid is not solved. Grid: \n%s",
				table)
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
		solved, _ := SolveBf(table)
		if !isSolved(solved) {
			t.Errorf("TestSolve was incorrect, grid is not solved. Grid: \n%s",
				table)
		}
	}
}

func TestSolveHuman(t *testing.T) {
	tables := []struct {
		grid       *Grid
		difficulty Difficulty
		solvable   bool
	}{
		{Parse("__,__"), DifficultyEasy, false},
		{Parse("_f_,___,___"), DifficultyEasy, false},
		{Parse("_f_,___,___"), DifficultyHard, true},
		{Parse("_f_,_f_,___"), DifficultyEasy, true},
	}

	for _, table := range tables {
		solution, _ := SolveHuman(table.grid, table.difficulty)
		if isSolved(solution) && !table.solvable {
			t.Errorf("TestSolveHuman was incorrect, grid is solved, but shouldn't be (difficulty: %v). Grid: \n%v",
				table.difficulty, table.grid)
		}
		if !isSolved(solution) && table.solvable {
			t.Errorf("TestSolveHuman was incorrect, grid is not solved (difficulty: %v). Grid: \n%v",
				table.difficulty, table.grid)
		}
	}
}

// history:
// - 36710 ns/op
// - 34716 ns/op (after switching to partially validation)
func BenchmarkSolve(b *testing.B) {
	g := Parse("_f_,_f_,_f_")
	for i := 0; i < b.N; i++ {
		solved, _ := Solve(g)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
		}
	}
}

func BenchmarkSolveDk(b *testing.B) {
	g := Parse("_f_,_f_,_f_")
	for i := 0; i < b.N; i++ {
		solved, _ := SolveDk(g)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
		}
	}
}

func BenchmarkSolveBf(b *testing.B) {
	g := Parse("_f_,_f_,_f_")
	for i := 0; i < b.N; i++ {
		solved, _ := SolveBf(g)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
		}
	}
}
