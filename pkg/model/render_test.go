package model

import (
	"testing"
)

func TestRender(t *testing.T) {
	g := Parse("___,df_,__x")

	_ = Render(g, 0)
}
