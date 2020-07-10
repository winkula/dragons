package model

type knowledge struct {
	size int
	pv   []uint8 // possible values
}

func newKnowledge(w *World) *knowledge {
	k := &knowledge{}
	k.size = w.Size()
	k.pv = make([]uint8, k.size)
	all := uint8(1<<SquareDragon | 1<<SquareFire | 1<<SquareEmpty)
	for i := range w.Squares {
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

func (k *knowledge) getOptions(w *World, i int) []Square {
	var res []Square
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
	perms []*World
}

func (k *knowledge) getPermutations(w *World, is []int) *permRes {
	result := &permRes{}
	permRecur(k, result, w, is, 0)
	return result
}

func permRecur(k *knowledge, result *permRes, w *World, indexes []int, i int) {
	if i >= len(indexes) {
		result.count++
		if ValidateWorld(w) {
			//fmt.Printf("- opt:\n%v\n", w)
			result.valid++
		}
		return // stop recursion
	}
	currentIndex := indexes[i]
	if w.GetSquareByIndex(currentIndex) == SquareUndefined {
		for _, v := range k.getOptions(w, currentIndex) {
			successor := w.Clone()
			successor.SetSquareByIndex(currentIndex, v)
			permRecur(k, result, successor, indexes, i+1)
		}
	} else {
		successor := w.Clone()
		permRecur(k, result, successor, indexes, i+1)
	}
}
