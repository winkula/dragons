package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

var validateCmd = flag.NewFlagSet("validate", flag.ExitOnError)

func init() {
	registerCommand("validate", validateCmd, func() {
		g := parse(validateCmd.Arg(0), true)
		validate(g)
	})
}

func validate(g *model.Grid) {
	valid := model.Validate(g)
	fmt.Println("Valid:", valid)
}
