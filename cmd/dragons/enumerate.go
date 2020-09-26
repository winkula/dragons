package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/winkula/dragons/pkg/model"
)

var enumerateCmd = flag.NewFlagSet("enumerate", flag.ExitOnError)
var enuArgMost = enumerateCmd.Bool("most", false, "only print the most interesting one")

func init() {
	registerCommand("enumerate", enumerateCmd, func() {
		g := parse(enumerateCmd.Arg(0), true)
		enumerate(g, *enuArgMost)
	})
}

func enumerate(g *model.Grid, most bool) {
	sucs := model.Enumerate(g)
	sort.Slice(sucs, func(i, j int) bool {
		return sucs[i].Interestingness() < sucs[j].Interestingness()
	})

	printThem := func(gs []*model.Grid) {
		for _, s := range gs {
			fmt.Println(s)
			fmt.Println("Interestingness:", s.Interestingness())
		}
	}

	if most {
		printThem(sucs[len(sucs)-1:])
	} else {
		printThem(sucs)
	}

	fmt.Printf("%v possible grids found.\n", len(sucs))
}
