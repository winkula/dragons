package model

import "testing"

// TODO
func TestSolve(t *testing.T) {
	tables := []*World{
		ParseWorld("dx,_x"),
		ParseWorld("dx,__"),
		ParseWorld("_f_,_f_,_f_"),
		//ParseWorld("d__,___,__d"),
	}

	for _, table := range tables {
		solved := Solve(table)
		if solved == nil || !isSolved(solved) {
			t.Errorf("TestSolve was incorrect, world is not solved. World: \n%s",
				table)
		}
	}
}
