package model

import (
	"testing"
	"time"
)

var duration = time.Second / 2

func TestGenerate(t *testing.T) {
	grid := Generate(3, 3, duration)

	if !IsDistinct(grid) {
		t.Errorf("TestGenerate was incorrect, grid is not distinct. Grid: \n%s", grid)
	}
}

func TestObfuscate(t *testing.T) {
	grid := Obfuscate(Parse("xfx,dfd,xfx"), DifficultyEasy, duration)

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

	Obfuscate(Parse("___,___,___"), DifficultyEasy, duration)
}

func BenchmarkGenerate(b *testing.B) {
	size := 3
	for i := 0; i < b.N; i++ {
		Generate(size, size, 50*time.Millisecond)
	}
}

// history:
// - 16471082 ns/op
//
func BenchmarkObfuscate(b *testing.B) {
	difficulty := Difficulty(DifficultyEasy)
	template := Parse("_f_,___,___")
	for i := 0; i < b.N; i++ {
		Obfuscate(template, difficulty, 5)
	}
}
