package renderers

import (
	"fmt"
	"image/color"
	"os"

	"github.com/tdewolff/canvas"
	"github.com/tdewolff/canvas/renderers/pdf"
	"github.com/winkula/dragons/pkg/model"
)

var size = 5 // 5mm
var sizeFactor = float64(size)
var padding = 1 // 1mm
var gridLine = 0.05
var gridBorder = 0.3
var gridColor = canvas.Black
var symbolLine = 0.15

func RenderPdf(g *model.Grid, border bool, filename string) {
	f, err := os.Create(fmt.Sprintf("%v.pdf", filename))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	p := pdf.New(f, float64(g.Width*size+2*padding), float64(g.Height*size+2*padding), nil)

	c := canvas.New(float64(g.Width*size+2*padding), float64(g.Height*size+2*padding))
	ctx := canvas.NewContext(c)
	ctx.SetFillColor(canvas.Transparent)
	ctx.SetStrokeCapper(canvas.SquareCap)
	ctx.SetStrokeJoiner(canvas.RoundJoin)

	drawGrid(ctx, g, border)
	drawSymbols(ctx, g)

	c.Render(p)
	p.Close()
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
		case model.SquareNoDragon:
			drawPoint(ctx, x, yNorm)
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
	triangleSize := 0.75
	sqrt3 := 1.73205080757
	triangleHeight := sqrt3 / 2.0 * triangleSize

	polyline := &canvas.Polyline{}
	polyline.Add((1.0-triangleSize)/2.0*sizeFactor, (1.0-triangleHeight)/2.0*sizeFactor)
	polyline.Add(0.5*sizeFactor, (1.0-(1.0-triangleHeight)/2.0)*sizeFactor)
	polyline.Add((1.0-(1.0-triangleSize)/2.0)*sizeFactor, (1.0-triangleHeight)/2.0*sizeFactor)
	polyline.Add((1.0-triangleSize)/2.0*sizeFactor, (1.0-triangleHeight)/2.0*sizeFactor)
	ctx.SetFillColor(col)
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(symbolLine)

	ctx.DrawPath(float64(size*x+padding), float64(size*y+padding), polyline.ToPath())
}

func drawLine(ctx *canvas.Context, x int, y int) {
	lineSize := 0.4
	polyline := &canvas.Polyline{}
	polyline.Add((1.0-lineSize)/2.0*sizeFactor, 0.5*sizeFactor)
	polyline.Add((1.0-(1.0-lineSize)/2.0)*sizeFactor, 0.5*sizeFactor)
	ctx.SetFillColor(canvas.Transparent)
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(symbolLine)

	ctx.DrawPath(float64(size*x+padding), float64(size*y+padding), polyline.ToPath())
}

func drawPoint(ctx *canvas.Context, x int, y int) {
	pointSize := 0.02 * sizeFactor
	ctx.SetFillColor(canvas.Black)
	ctx.SetStrokeColor(canvas.Black)
	ctx.SetStrokeWidth(symbolLine)

	circle := canvas.Circle(pointSize).Translate(0.5*sizeFactor, 0.5*sizeFactor)
	ctx.DrawPath(float64(size*x+padding), float64(size*y+padding), circle)
}
