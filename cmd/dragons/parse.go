package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

func init() {
	cmd := flag.NewFlagSet("parse", flag.ExitOnError)

	registerCommand("parse", cmd, func() {
		parse(cmd.Arg(0), true)
	})
}

func parse(s string, print bool) *model.Grid {
	g := model.Parse(s)
	if print {
		fmt.Println("Grid:")
		fmt.Println(g)
	}
	return g
}
