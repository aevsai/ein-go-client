package objects

import "github.com/fogleman/gg"

type Renderable interface {
	Render(ctx *gg.Context)
}

type BaseObject struct {
	X     float64
	Y     float64
	Color string
}

type Circle struct {
	BaseObject
	Radius float64
}

func (c Circle) Render(ctx *gg.Context) {
	ctx.SetHexColor(c.Color)
	ctx.DrawCircle(c.X, c.Y, c.Radius)
	ctx.Fill()
}

type Rectangle struct {
	BaseObject
	Width  float64
	Height float64
}

func (r Rectangle) Render(ctx *gg.Context) {
	ctx.SetHexColor(r.Color)
	ctx.DrawRectangle(r.X, r.Y, r.Width, r.Height)
	ctx.Fill()
}

type Line struct {
	BaseObject
	X2 float64
	Y2 float64
}

func (l Line) Render(ctx *gg.Context) {
	ctx.SetHexColor(l.Color)
	ctx.DrawLine(l.X, l.Y, l.X2, l.Y2)
	ctx.Stroke()
}

type Arc struct {
	BaseObject
	Radius float64
	Start  float64
	End    float64
}

func (a Arc) Render(ctx *gg.Context) {
	ctx.SetHexColor(a.Color)
	ctx.DrawArc(a.X, a.Y, a.Radius, a.Start, a.End)
	ctx.Fill()
}

type Polygon struct {
	BaseObject
	N        int
	R        float64
	Rotation float64
}

func (p Polygon) Render(ctx *gg.Context) {
	ctx.SetHexColor(p.Color)
	ctx.DrawRegularPolygon(p.N, p.X, p.Y, p.R, p.Rotation)
	ctx.Fill()
}
