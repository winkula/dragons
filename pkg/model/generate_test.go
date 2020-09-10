package model

import "testing"

func TestGenerate(t *testing.T) {
	grid := Generate(3, 3, .5)

	if !IsDistinct(grid) {
		t.Errorf("TestGenerate was incorrect, grid is not distinct. Grid: \n%s", grid)
	}
}

func TestGenerateFrom(t *testing.T) {
	grid := GenerateFrom(Parse("xfx,dfd,xfx"), DifficultyEasy, .5)

	if !IsDistinct(grid) {
		t.Errorf("TestGenerateFrom was incorrect, grid is not distinct. Grid: \n%s", grid)
	}
}

func TestGenerateFrom_WithUndefinedGrid_ShouldPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("TestGenerateFrom_WithUndefinedGrid_ShouldPanic was incorrect, the code did not panic")
		}
	}()

	GenerateFrom(Parse("___,___,___"), DifficultyEasy, .5)
}

func BenchmarkGenerate(b *testing.B) {
	size := 3
	for i := 0; i < b.N; i++ {
		Generate(size, size, 10.0)
	}
}

// history:
// - 16471082 ns/op
//
func BenchmarkGenerateFrom(b *testing.B) {
	difficulty := Difficulty(DifficultyEasy)
	template := Parse("_f_,___,___")
	for i := 0; i < b.N; i++ {
		GenerateFrom(template, difficulty, 5)
	}
}
