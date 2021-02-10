package renderers

import (
	"image/color"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/pdf"
	"github.com/tdewolff/canvas/svg"
	"github.com/winkula/dragons/pkg/model"
)

var size = 5 // 5mm
var sizeFactor = float64(size)
var padding = 1 // 1mm
var gridLine = 0.05
var gridBorder = 0.25
var gridColor = canvas.Black
var symbolLine = 0.15

func RenderPdf(g *model.Grid, border bool) {
	c := canvas.New(float64(g.Width*size+2*padding), float64(g.Height*size+2*padding))
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.Transparent)
	ctx.SetStrokeCapper(canvas.SquareCap)
	ctx.SetStrokeJoiner(canvas.RoundJoin)

	drawGrid(ctx, g, border)
	drawSymbols(ctx, g)

	c.WriteFile("out.pdf", pdf.Writer)
	c.WriteFile("out.svg", svg.Writer)
}

func drawGrid(ctx *canvas.Context, g *model.Grid, border bool) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			ctx.SetFillColor(canvas.Transparent)
			ctx.SetStrokeColor(gridColor)
			ctx.SetStrokeWidth(gridLine)

			drawRect(ctx, sizeFactor*float64(x), sizeFactor*float64(y), sizeFactor, sizeFactor, gridLine, gridColor, canvas.White)
		}
	}

	if border {
		drawRect(ctx, 0.0, 0.0, sizeFactor*float64(g.Width), sizeFactor*float64(g.Height), gridBorder, gridColor, canvas.Transparent)
	}
}

func drawSymbols(ctx *canvas.Context, g *model.Grid) {
	for i, square := range g.Squares {
		x, y := g.Coords(i)

		yNorm := g.Height - y - 1

		switch square {
		case model.SquareDragon:
			drawTriangle(ctx, x, yNorm, canvas.Black)
		case model.SquareFire:
			drawTriangle(ctx, x, yNorm, canvas.Transparent)
		case model.SquareAir:
			drawLine(ctx, x, yNorm)
		}
	}
}

func drawRect(ctx *canvas.Context, x float64, y float64, width float64, height float64, strokeWidth float64, strokeColor color.RGBA, fillColor color.RGBA) {
	polyline := &canvas.Polyline{}
	polyline.Add(0.0, 0.0)
	polyline.Add(width, 0.0)
	polyline.Add(width, height)
	polyline.Add(0.0, height)
	polyline.Add(0.0, 0.0)
	ctx.SetFillColor(fillColor)
	ctx.SetStrokeColor(strokeColor)
	ctx.SetStrokeWidth(strokeWidth)

	ctx.DrawPath(x+float64(padding), y+float64(padding), polyline.ToPath())
}

func drawTriangle(ctx *canvas.Context, x int, y int, col color.RGBA) {
	polyline := &canvas.Polyline{}
	polyline.Add(0.2*sizeFactor, 0.26*sizeFactor)
	polyline.Add(0.5*sizeFactor, 0.78*sizeFactor)
	polyline.Add(0.8*sizeFactor, 0.26*sizeFactor)
	polyline.Add(0.2*sizeFactor, 0.26*sizeFactor)
	ctx.SetFillColor(col)
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(symbolLine)

	ctx.DrawPath(float64(size*x+padding), float64(size*y+padding), polyline.ToPath())
}

func drawLine(ctx *canvas.Context, x int, y int) {
	polyline := &canvas.Polyline{}
	polyline.Add(0.3*sizeFactor, 0.5*sizeFactor)
	polyline.Add(0.7*sizeFactor, 0.5*sizeFactor)
	ctx.SetFillColor(canvas.Transparent)
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(symbolLine)

	ctx.DrawPath(float64(size*x+padding), float64(size*y+padding), polyline.ToPath())
}
