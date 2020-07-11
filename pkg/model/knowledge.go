package model

type knowledge struct {
	size int
	pv   []uint8 // possible values
}

func newKnowledge(g *Grid) *knowledge {
	k := &knowledge{}
	k.size = g.Size()
	k.pv = make([]uint8, k.size)
	all := uint8(1<<SquareDragon | 1<<SquareFire | 1<<SquareEmpty)
	for i := range g.Squares {
		k.pv[i] = all
	}
	return k
}

func (k *knowledge) squareCannotBe(i int, v Square) {
	if i < 0 || i >= k.size {
		return
	}
	k.pv[i] &= ^(1 << v)
}

func (k *knowledge) squaresCannotBe(is []int, v Square) {
	for _, i := range is {
		k.squareCannotBe(i, v)
	}
}

func (k *knowledge) canSquareBe(i int, v Square) bool {
	if i < 0 || i >= k.size {
		return false
	}
	return k.pv[i]&(1<<v) > 0
}

func (k *knowledge) getOptions(g *Grid, i int) []Square {
	res := make([]Square, 0, len(options))
	for _, o := range options {
		if k.canSquareBe(i, o) {
			res = append(res, o)
		}
	}
	return res
}

type permRes struct {
	count int
	valid int
	perms []*Grid
}

func (k *knowledge) getPermutations(g *Grid, is []int) *permRes {
	result := &permRes{}
	permRecur(k, result, g, is, 0)
	return result
}

func permRecur(k *knowledge, result *permRes, g *Grid, indexes []int, i int) {
	if i >= len(indexes) {
		result.count++
		if Validate(g) {
			result.valid++
		}
		return // stop recursion
	}
	currentIndex := indexes[i]
	if g.Squarei(currentIndex) == SquareUndefined {
		for _, v := range k.getOptions(g, currentIndex) {
			suc := g.Clone()
			suc.SetSquarei(currentIndex, v)
			permRecur(k, result, suc, indexes, i+1)
		}
	} else {
		suc := g.Clone()
		permRecur(k, result, suc, indexes, i+1)
	}
}
