package model

import "testing"

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
