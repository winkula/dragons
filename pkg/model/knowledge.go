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

func (k *knowledge) possibleCount(is []int, v Square) {

}
