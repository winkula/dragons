package model

type solveResult int

// TODO: solve rules should only return true when actually modifications were done
// or when new knowledge was gained
const (
	solveResultNone            = 0
	solveResultUpdateKnowledge = 1
	solveResultFillSquare      = 2
)

type solveRule (func(*Grid, int, *knowledge) solveResult)

// rule 1
// if a dragon is set, no neighbour square can be dragons too
var rule1 = func(g *Grid, i int, k *knowledge) solveResult {
	if g.Squarei(i) == SquareDragon {
		k.squaresCannotBe(g.NeighborIndicesi(i, false), SquareDragon)
		debug("-> [ar] neighbour squares of dragon cannot be dragons")
		return solveResultUpdateKnowledge
	}
	return solveResultNone
}

// rule 2
// if a undefined square is surrounded by more than one dragon, there must be fire
var rule2 = func(g *Grid, i int, k *knowledge) solveResult {
	if g.Squarei(i) == SquareUndefined {
		dragons := g.NeighborCount8(i, SquareDragon)
		if dragons > 1 {
			g.SetSquarei(i, SquareFire)
			debug("-> [ar] set fire if more than 1 dragon around square")
			debug(g)
			return solveResultFillSquare
		}
	}
	return solveResultNone
}

// rule 2 (derived)
// if a air square is set, there can maximum be one dragon around it
// otherwise the square would be fire, not air
var rule2derived = func(g *Grid, i int, k *knowledge) solveResult {
	if g.Squarei(i) == SquareAir {
		dragons := g.NeighborCount8(i, SquareDragon)
		if dragons == 1 {
			k.squaresCannotBe(g.NeighborIndicesi(i, false), SquareDragon)
			debug("-> [ar] if a air square is set, there can maximum be one dragon around it")
			return solveResultUpdateKnowledge
		}
	}
	return solveResultNone
}

// rule 2 (inverted)
// if a square is fire, there must be at least 2 dragons around it
var rule2inverted = func(g *Grid, i int, k *knowledge) solveResult {
	if g.Squarei(i) == SquareFire {
		dragons := g.NeighborCount8(i, SquareDragon)
		undef := g.NeighborCount8(i, SquareUndefined)
		if dragons < 2 && dragons+undef == 2 {
			for _, ni := range g.NeighborIndicesi(i, false) {
				if g.Squarei(ni) == SquareUndefined {
					g.SetSquarei(ni, SquareDragon)
				}
			}
			debug("-> [ar] set dragons in neighbour squares if only one solution")
			debugGrid(g, i)
			return solveResultFillSquare
		}
	}
	return solveResultNone
}

// rule 2 (inverted)
// if a undefined square cannot be surrounded be at least 2 dragons, there can not be fire
//
// TODO: check if we want this rule to be part of the solver
// the problem is, that if we use this rule, one must be able to write down the hint "can not be fire" on the paper
var rule2invertedSpecial = func(g *Grid, i int, k *knowledge) solveResult {
	if g.Squarei(i) == SquareUndefined {
		dragons := g.NeighborCount8(i, SquareDragon)
		undef := g.NeighborCount8(i, SquareUndefined)
		if dragons+undef < 2 {
			k.squareCannotBe(i, SquareFire)
			debug("-> [ar] square cannot be fire if not at least 2 dragons could be around it")
			return solveResultUpdateKnowledge
		}
	}
	return solveResultNone
}

// rule 3
// if a dragon is set, at least two adjacent quares must be air
var rule3 = func(g *Grid, i int, k *knowledge) solveResult {
	if g.Squarei(i) == SquareDragon {
		air := g.NeighborCount4(i, SquareAir)
		undef := g.NeighborCount4(i, SquareUndefined)
		if air < 2 && air+undef == 2 {
			for _, ni := range g.NeighborIndicesi(i, true) {
				if g.Squarei(ni) == SquareUndefined {
					g.SetSquarei(ni, SquareAir)
				}
			}
			debug("-> [ar] fill adjacent squares to air")
			debugGrid(g, i)
			return solveResultFillSquare
		}
	}
	return solveResultNone
}

// rule 3 (inverted)
// a square can not be a dragon if there are not 2 air possible in the adjacent squares
var rule3Inverted = func(g *Grid, i int, k *knowledge) solveResult {
	if g.Squarei(i) == SquareUndefined {
		air := g.NeighborCount4(i, SquareAir)
		undef := g.NeighborCount4(i, SquareUndefined)
		if air+undef < 2 {
			k.squareCannotBe(i, SquareDragon)
			debug("-> [ar] square connot be a dragon because there could not be 2 adjacent squares of air")
			return solveResultUpdateKnowledge
		}
	}
	return solveResultNone
}

// no rule - just used to apply knowledge to the grid
// if a square is undefined but there is only one value possible (according to the knowledge db)
// set the squares value
var applyKnowledge = func(g *Grid, i int, k *knowledge) solveResult {
	if g.Squarei(i) == SquareUndefined && k.optionsCount(i) == 1 {
		val := k.onlyPossibleValue(i)
		g.SetSquarei(i, val)
		debug("-> [ar] set the only posible value of a square")
		debugGrid(g, i)
		return solveResultFillSquare
	}
	return solveResultNone
}

var solveRulesEasy = []solveRule{
	rule1,
	rule2,
	rule2derived,
	rule2inverted,
	//rule2invertedSpecial, // excluded rule
	rule3,
	rule3Inverted,
	applyKnowledge,
}

var solveRulesMedium = []solveRule{
	rule1,
	rule2,
	rule2derived,
	rule2inverted,
	rule2invertedSpecial,
	rule3,
	rule3Inverted,
	applyKnowledge,
}

var solveTechniques = map[Difficulty][]solveRule{
	DifficultyEasy:   solveRulesEasy,
	DifficultyMedium: solveRulesMedium,
	DifficultyHard:   solveRulesMedium,
}
