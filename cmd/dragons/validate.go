package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

func init() {
	cmd := flag.NewFlagSet("validate", flag.ExitOnError)

	registerCommand("validate", cmd, func() {
		g := parse(cmd.Arg(0), true)
		validate(g)
	})
}

func validate(g *model.Grid) {
	valid := model.Validate(g)
	fmt.Println("Valid:", valid)
}
