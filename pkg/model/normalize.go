package model

// Normalize normalizes a grid to prevent duplicates (rotated and mirrored grids).
func (g *Grid) Normalize() *Grid {
	rLeft := g.Rotate(true)
	rRight := g.Rotate(false)
	mHor := g.Mirror(true)
	mVert := g.Mirror(false)
	mBoth := mHor.Mirror(false)
	return withLowestCode([]*Grid{g, rLeft, rRight, mHor, mVert, mBoth})
}

func withLowestCode(grids []*Grid) *Grid {
	code := grids[0].ID()
	best := grids[0]
	for _, g := range grids {
		c := g.ID()
		if c.Cmp(code) == -1 {
			code = c
			best = g
		}
	}
	return best
}

// Rotate creates a rotated copy of a grid.
func (g *Grid) Rotate(left bool) *Grid {
	n := New(g.Height, g.Width)
	if left {
		for i, s := range g.Squares {
			x, y := g.Coords(i)
			n.SetSquare(y, g.Width-x-1, s)
		}
	} else {
		for i, s := range g.Squares {
			x, y := g.Coords(i)
			n.SetSquare(g.Height-y-1, x, s)
		}
	}
	return n
}

// Mirror creates a mirrored copy of a grid.
func (g *Grid) Mirror(horizontal bool) *Grid {
	n := New(g.Width, g.Height)
	if horizontal {
		for i, s := range g.Squares {
			x, y := g.Coords(i)
			n.SetSquare(g.Width-x-1, y, s)
		}
	} else {
		for i, s := range g.Squares {
			x, y := g.Coords(i)
			n.SetSquare(x, g.Height-y-1, s)
		}
	}
	return n
}
