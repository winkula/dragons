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

func TestSolveDk(t *testing.T) {
	tables := []struct {
		grid       *Grid
		difficulty Difficulty
		solvable   bool
	}{
		{Parse("dx,_x"), DifficultyMedium, true},
		{Parse("dx,__"), DifficultyMedium, true},
		{Parse("_f_,_f_,_f_"), DifficultyMedium, true},
		{Parse("d__,___,__d"), DifficultyMedium, true},

		{Parse("_f_,_f_,_f_"), DifficultyEasy, false},   // not solvable using easy rule set
		{Parse("_f_,___,___"), DifficultyMedium, false}, // not solvable
		{Parse("___,_f_,___"), DifficultyMedium, false}, // no distinct solution possible
	}

	for _, table := range tables {
		solved := SolveDk(table.grid, table.difficulty)

		if table.solvable {
			if solved == nil || !isSolved(solved) {
				t.Errorf("TestSolve was incorrect, grid is not solved. Grid: \n%s",
					table.grid)
			}
		} else {
			if solved != nil {
				t.Errorf("TestSolve was incorrect, grid is solved. Grid: \n%s",
					table.grid)
			}
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
		solution := SolveHuman(table.grid, table.difficulty)
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
		solved := Solve(g)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
		}
	}
}

func BenchmarkSolveDk(b *testing.B) {
	g := Parse("_f_,_f_,_f_")
	for i := 0; i < b.N; i++ {
		solved := SolveDk(g, DifficultyEasy)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
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
