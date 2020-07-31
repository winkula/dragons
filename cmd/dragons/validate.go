package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
)

var validateCmd = flag.NewFlagSet("validate", flag.ExitOnError)

func validate(g *model.Grid) {
	valid := model.Validate(g)
	fmt.Println("Valid:", valid)
}
