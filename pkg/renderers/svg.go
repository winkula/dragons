package renderers

import (
	"bufio"
	"fmt"
	"os"
	"text/template"

	"github.com/winkula/dragons/pkg/model"
)

var templ = `
{{- define "fire" }}
	<g class="fire" transform="translate({{ .X }},{{ .Y }})">
		<!--<path d="M 50 15 L 85 80 L 15 80 L 50 15 Z"></path>-->
		<path d="M 50 22 L 80 74 L 20 74 L 50 22 Z"></path>
	</g>
{{ end -}}
{{- define "dragon" }}
	<g class="dragon" transform="translate({{ .X }},{{ .Y }}) translate(50, 50) scale(0.9) translate(-50,-50)">
		<path d="M 9.345 54.212 C 30.721 31.007 56.325 25.154 89.64 25.996 C 85.485 80.384 23.579 90.086 9.345 54.212 Z"></path>
		<circle cx="44.706" cy="29.788" r="22.912" transform="matrix(0.209362, 0, 0, 0.85314, 40.429417, 24.468189)"></circle>
	</g>
{{ end -}}
{{- define "air" }}
	<g class="air" transform="translate({{ .X }},{{ .Y }})">
		<line x1="20" y1="50" x2="80" y2="50"></line>
	</g>
{{ end -}}
{{- define "value" }}
	{{ if eq .Template "fire" }}
		{{ template "fire" . }}
	{{ else if eq .Template "dragon" }}
		{{ template "dragon" . }}
	{{ else if eq .Template "air" }}
		{{ template "air" . }}	
	{{ end }}
{{ end -}}
{{- define "square" }}
	<rect class="square" x="{{ .X }}" y="{{ .Y }}" width="{{ .Size }}" height="{{ .Size }}" shape-rendering="crispEdges" />
{{ end -}}

<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<svg
 xmlns="http://www.w3.org/2000/svg"
 width="{{ .Width }}"
 height="{{ .Height }}"
 viewBox="0 0 {{ .Width }} {{ .Height }}"
>
	<style type="text/css">
		<![CDATA[
		* {
			vector-effect: non-scaling-stroke;
		}
		.fire path {
			fill: #fff;
			stroke: #000;
			stroke-width: 3;
		}
		.square {
			fill: #fff;
			stroke: #000;
			stroke-width: 1;
		}
		.dragon path {
			stroke: #000;
			stroke-width: 8;
			fill: #fff;
		}
		.dragon circle {
			stroke: #000;
			stroke-width: 8;
			fill: #000;
		}
		.dragon circle {
			fill: #000;
		}
		.air line {
			stroke: #000;
			stroke-width: 3;
		}
		.outline {
			stroke: #000;
			stroke-width: 7;
			fill: none;
		}	
		]]>
	</style> 


	{{ range $square := .Squares }}
		{{ template "square" $square }}
		{{ template "value" $square }}
	{{end}}
	
	{{ if eq .Border true }}
		<rect class="outline" x="0" y="0" width="{{ .Width }}" height="{{ .Height }}" shape-rendering="crispEdges" />
	{{ end }}
</svg> 
`

type svgGrid struct {
	Squares []svgSquare
	Width   string
	Height  string
	Border  bool
}

type svgSquare struct {
	X        int
	Y        int
	Fill     string
	Template string
	Size     int
	SizeHalf int
}

// RenderSvg prints a grid in SVG format
func RenderSvg(g *model.Grid, border bool) {
	gridSize := 100
	grid := svgGrid{
		Squares: make([]svgSquare, 0, g.Width*g.Height),
		Width:   fmt.Sprintf("%v", g.Width*gridSize),
		Height:  fmt.Sprintf("%v", g.Height*gridSize),
		Border:  border,
	}
	for i, square := range g.Squares {
		x, y := g.Coords(i)
		cell := svgSquare{
			X:        x * gridSize,
			Y:        y * gridSize,
			Fill:     []string{"", "", "white", "red", "black"}[square],
			Template: []string{"", "", "air", "fire", "dragon"}[square],
			Size:     gridSize,
			SizeHalf: gridSize / 2,
		}
		grid.Squares = append(grid.Squares, cell)
	}

	saveFile("test.svg", grid)
}

func saveFile(filename string, data interface{}) {
	f, _ := os.Create(filename)
	defer f.Close()

	w := bufio.NewWriter(f)
	t, _ := template.New("").Parse(templ)
	t.Execute(w, data)

	w.Flush()
}
