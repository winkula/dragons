package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/winkula/dragons/pkg/model"
)

func init() {
	cmd := flag.NewFlagSet("enumerate", flag.ExitOnError)
	most := cmd.Bool("most", false, "only print the most interesting one")
	interestingness := cmd.Int("i", 0, "minimum interestingness")

	registerCommand("enumerate", cmd, func() {
		g := parse(cmd.Arg(0), true)
		enumerate(g, *most, *interestingness)
	})
}

func enumerate(g *model.Grid, most bool, interestingness int) {
	minInterestingness := func(g *model.Grid) bool {
		return g.SolutionRating() >= float64(interestingness)
	}

	grids := model.EnumerateFilter(g, minInterestingness)

	sortByInterestingness := func(i, j int) bool {
		return grids[i].SolutionRating() < grids[j].SolutionRating()
	}

	sort.Slice(grids, sortByInterestingness)

	printThem := func(gs []*model.Grid) {
		for _, s := range gs {
			fmt.Println(s)
			fmt.Println("Interestingness:", s.SolutionRating())
		}
	}

	if most {
		printThem(grids[len(grids)-1:])
	} else {
		printThem(grids)
	}

	fmt.Printf("%v possible grids found.\n", len(grids))
}
