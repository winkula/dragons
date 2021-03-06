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

func TestSolveTechnically(t *testing.T) {
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
		solved := SolveTechnically(table.grid, table.difficulty)

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

func TestSolveBruteForce(t *testing.T) {
	tables := []*Grid{
		Parse("dx,_x"),
		Parse("dx,__"),
		Parse("_f_,_f_,_f_"),
		Parse("_f_,___,___"),
		Parse("d__,___,__d"),
	}

	for _, table := range tables {
		solved := SolveBruteForce(table)
		if !isSolved(solved) {
			t.Errorf("SolveBruteForce was incorrect, grid is not solved. Grid: \n%s",
				table)
		}
	}
}

func TestSolveBruteForceWithUnsolvablePuzzle(t *testing.T) {
	tables := []*Grid{
		Parse("d_,_d"),
		Parse("_f_,__f,_f_"),
	}

	for _, table := range tables {
		solved := SolveBruteForce(table)
		if solved != nil {
			t.Errorf("SolveBruteForce was incorrect, puzzle should not be solvable.")
		}
	}
}

func TestSolveIterative(t *testing.T) {
	tables := []struct {
		grid       *Grid
		difficulty Difficulty
		solvable   bool
	}{
		{Parse("__,__"), DifficultyEasy, false},
		{Parse("_f_,___,___"), DifficultyEasy, false},
		{Parse("_f_,___,___"), DifficultyBrutal, true},
		{Parse("_f_,_f_,___"), DifficultyEasy, true},
		{Parse("_f_,_f_,___"), DifficultyMedium, true},
	}

	for _, table := range tables {
		solution := SolveIterative(table.grid, table.difficulty)
		if isSolved(solution) && !table.solvable {
			t.Errorf("TestSolveIterative was incorrect, grid is solved, but shouldn't be (difficulty: %v). Grid: \n%v",
				table.difficulty, table.grid)
		}
		if !isSolved(solution) && table.solvable {
			t.Errorf("TestSolveIterative was incorrect, grid is not solved (difficulty: %v). Grid: \n%v",
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

func BenchmarkSolveTechnically(b *testing.B) {
	g := Parse("_f_,_f_,_f_")
	for i := 0; i < b.N; i++ {
		solved := SolveTechnically(g, DifficultyEasy)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
		}
	}
}

func BenchmarkSolveBruteForce(b *testing.B) {
	g := Parse("_f_,_f_,_f_")
	for i := 0; i < b.N; i++ {
		solved := SolveBruteForce(g)
		if solved == nil || !isSolved(solved) {
			b.Errorf("BenchmarkSolve was incorrect, grid is not solved. Grid: \n%s", g)
		}
	}
}
