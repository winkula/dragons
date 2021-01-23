package main

import (
	"flag"

	"github.com/winkula/dragons/pkg/model"
	"github.com/winkula/dragons/pkg/renderers"
)

func init() {
	cmd := flag.NewFlagSet("render", flag.ExitOnError)
	noOutline := cmd.Bool("no-outline", false, "render no outline around the grid")

	registerCommand("render", cmd, func() {
		render(cmd.Arg(0), *noOutline)
	})
}

func render(s string, noOutline bool) {
	g := model.Parse(s)
	renderers.RenderSvg(g, !noOutline)
}
