package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

var parseCmd = flag.NewFlagSet("parse", flag.ExitOnError)

func init() {
	registerCommand("parse", parseCmd, func() {
		parse(parseCmd.Arg(0), true)
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
