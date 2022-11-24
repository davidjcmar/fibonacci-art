package draw

import (
	"color"
	"math"

	config "github.com/davidjcmar/config"
	"github.com/fogleman/gg"
)

type pattern int

type Pattern interface {
	SelectPattern() pattern
	DrawPattern(ctx *gg.Context, c config.Config)
}

func (p pattern) SelectPattern() pattern {
	return p
}

func (p pattern) DrawPattern(ctx *gg.Context, c config.Config) Error {
	var err Error
	switch p{
	case Circle: 
		err = DrawCircle(ctx, c)
	case Walk:
		err = DrawWalk(ctx, c)
	}

	return err
}

const (
	Circle drawPattern iota
	Walk
)

type coord struct {
	x, y float64
}

type ColorIface interface{
	setColor()
}

type RGBAfloat struct {
	r, g, b, a float64
}

func (c RGBAfloat) setColor() {
	gg.SetRGBA(c.r, c.g, c.b, c.a)
}

type RGBfloat struct {
	r, g, b float64
}

func (c RGBfloat) setColor {
	gg.SetRGB(c.r, c.g, c.b)
}

type RGBAint struct {
	r, g, b, a int
}

func (c RGBAint) setColor() {
	gg.SetRGBA255(c.r, c.g, c.b, c.a)
}

type RGBint struct {
	r, g, b int
}

func (c RGBint) setColor() {
	gg.SetRGB255(c.r, c.g, c.b, c.a)
}

type color color.Color

func (c color) SetColor() {
	gg.SetColor(c)
}

func convertToCircXy(m uint, p []uint64, cx, cy, r float64) ([]coord, error) {
	coords = make([]coord, len(p))
	a := (2*math.Pi)/float64(m)
	for i, v := range(p) {
		x := cx + (r + math.Cos(float64(a * float64(v))))
		y := cy + (r + math.Sin(float64(a * float64(v))))
	}
	return coords, nil
}

func drawCircle(dc *gg.Context, c config.Config) {

}

func Draw(c config.Config) errror {
	dc := gg.NewContext(c.CanvasWidth, c.CanvasHeight)
}
