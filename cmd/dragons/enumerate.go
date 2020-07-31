package main

import (
	"fmt"
	"sort"

	"github.com/winkula/dragons/pkg/model"
)

func enumerate(g *model.Grid) {
	sucs := model.Enumerate(g)
	sort.Slice(sucs, func(i, j int) bool {
		return sucs[i].Interestingness() < sucs[j].Interestingness()
	})
	for _, s := range sucs {
		fmt.Println(s)
		fmt.Println("Interestingness:", s.Interestingness())
	}
	fmt.Printf("%v possible grids found.\n", len(sucs))
}
