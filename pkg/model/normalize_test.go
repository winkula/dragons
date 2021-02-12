package model

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tables := []struct {
		grid     *Grid
		left     bool
		expected *Grid
	}{
		{Parse("df,_f,x_"), true, Parse("ff_,d_x")},
		{Parse("df,_f,x_"), false, Parse("x_d,_ff")},
	}

	for _, table := range tables {
		rotated := table.grid.Rotate(table.left)
		if !reflect.DeepEqual(rotated, table.expected) {
			t.Errorf("rotate was incorrect, got: \n%v, want: \n%v.",
				rotated,
				table.expected)
		}
	}
}

func TestMirror(t *testing.T) {
	tables := []struct {
		grid     *Grid
		left     bool
		expected *Grid
	}{
		{Parse("df,_f,x_"), true, Parse("fd,f_,_x")},
		{Parse("df,_f,x_"), false, Parse("x_,_f,df")},
	}

	for _, table := range tables {
		rotated := table.grid.Mirror(table.left)
		if !reflect.DeepEqual(rotated, table.expected) {
			t.Errorf("mirror was incorrect, got: \n%v, want: \n%v.",
				rotated,
				table.expected)
		}
	}
}

func TestNormalize(t *testing.T) {
	tables := []struct {
		grid     *Grid
		expected *Grid
	}{
		{Parse("__d"), Parse("d__")},
		{Parse("__,d_"), Parse("d_,__")},
		{Parse("fd,f_,_x"), Parse("df,_f,x_")},
		{Parse("df,_f,x_"), Parse("df,_f,x_")},
	}

	for _, table := range tables {
		normalized := table.grid.Normalize()
		if !reflect.DeepEqual(normalized, table.expected) {
			t.Errorf("Normalize was incorrect, got: \n%v, want: \n%v.",
				normalized,
				table.expected)
		}
	}
}
