package main

import (
	"flag"
	"fmt"

	"github.com/winkula/dragons/pkg/model"
	"github.com/winkula/dragons/pkg/renderers"
)

func init() {
	cmd := flag.NewFlagSet("render", flag.ExitOnError)
	filename := cmd.String("filename", "test", "filename for the generated output")
	noOutline := cmd.Bool("no-outline", false, "render no outline around the grid")

	registerCommand("render", cmd, func() {
		render(cmd.Arg(0), *filename, *noOutline)
	})
}

func render(s string, filename string, noOutline bool) {
	g := model.Parse(s)
	//renderers.RenderSvg(g, !noOutline, fmt.Sprintf("%v_old", filename))
	renderers.RenderPdf(g, !noOutline, filename)
}
