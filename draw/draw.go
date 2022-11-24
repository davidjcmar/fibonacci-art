package draw

import (
	"math"

	config "github.com/davidjcmar/fibonacci-art/config"
	"github.com/fogleman/gg"
)

// Type coord represents a point on a 2-D coordinate plane
type coord struct {
	x, y float64
}

// Function convertToCircXY a slice of uint64 to a slice of type coords along the perimeter of a circle given a modulo value m
func convertToCircXy(m uint, p []uint64, cx, cy, r float64) []coord {
	coords := make([]coord, len(p))
	a := (2*math.Pi)/float64(m)
	for i, v := range(p) {
		x := cx + (r * math.Cos(float64(a * float64(v))))
		y := cy + (r * math.Sin(float64(a * float64(v))))
		coords[i].x, coords[i].y = x, y
	}
	return coords
}

// Function drawBackground takes a config.Config struct and gg.Context and draws the background color for the generated image
func drawBackground(c config.Config, dc *gg.Context) {
	dc.SetRGB255(c.BackgroundColor[config.Red], c.BackgroundColor[config.Green], c.BackgroundColor[config.Blue])
	dc.DrawRectangle(0, 0, float64(c.CanvasWidth), float64(c.CanvasHeight))
	dc.Fill()
}

// Function drawCircle tages a config.Config stuct, a slice coords, and a gg.Context and draws the outer circle for the generated image
func drawCircle(c config.Config, coords []coord, dc *gg.Context) {
	// draw outline
	dc.SetRGB255(c.OutlineColor[config.Red], c.OutlineColor[config.Green], c.OutlineColor[config.Blue])
	dc.SetLineWidth(c.OutlineWidth)
	dc.DrawCircle(
		float64(c.CanvasWidth/2),
		float64(c.CanvasHeight/2),
		c.Radius + c.OutlineWidth,
	)
	dc.Fill()
	// replace inner circle with background color
	dc.SetRGB255(c.BackgroundColor[config.Red], c.BackgroundColor[config.Green], c.BackgroundColor[config.Blue])
	dc.SetLineWidth(0)
	dc.DrawCircle(
		float64(c.CanvasWidth/2),
		float64(c.CanvasHeight/2),
		c.Radius,
	)
	dc.Fill()
	
	// draw inscribed lines
	dc.SetRGB255(c.LineColor[config.Red], c.LineColor[config.Green], c.LineColor[config.Blue])
	dc.SetLineWidth(c.LineWidth)
	dc.RotateAbout(math.Pi/2, float64(c.CanvasWidth/2), float64(c.CanvasHeight/2))
	for i, _ := range(coords) {
		if i == len(coords) -1 {
			break
		}
		dc.DrawLine(coords[i].x, coords[i].y, coords[i+1].x, coords[i+1].y)
		dc.Stroke()
	}
}

// Function Draw takes a config.Config struct and a slice of uint64 pisano period values and orchestrates creating a gg.Context and calling other drawers
func Draw(c config.Config, p []uint64) {
	dc := gg.NewContext(c.CanvasWidth, c.CanvasHeight)
	drawBackground(c, dc)
	drawCircle(c, convertToCircXy(c.Modulo, p, float64(c.CanvasWidth/2), float64(c.CanvasHeight/2), c.Radius), dc)
	
	outputFile := c.OutputFile + "." + c.OutputType
	switch c.OutputType {
		case "png": 
			dc.SavePNG(outputFile)
		default:
			dc.SavePNG(outputFile)
	}
}
