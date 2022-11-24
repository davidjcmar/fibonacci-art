package main

import (
	"flag"
	"fmt"
	"os"

	config "github.com/davidjcmar/fibonacci-art/config"
	draw "github.com/davidjcmar/fibonacci-art/draw"
	period "github.com/davidjcmar/fibonacci-art/period"
)

// Function initArgs parses command line arguments and initializes the config.Config struct with the values
func initArgs() *config.Config {
	mod := flag.Uint("mod", 7, "Modulus for pisano period")
	cw := flag.Int("width", 250, "Canvas width")
	ch := flag.Int("height", 250, "Canvas height")
	r := flag.Float64("radius", 120, "Circle radius")
	lw := flag.Float64("line", 2, "Line width")
	ow := flag.Float64("outline", 10, "Outline width")
	lineRed := flag.Int("red", 100, "Line red color component (0-255)")
	lineGreen := flag.Int("green", 100, "Line green component (0-255)")
	lineBlue := flag.Int("blue", 100, "Line blue component (0-255)")
	outlineRed := flag.Int("ored", 25, "Outline red color compoment (0-255)")
	outlineGreen := flag.Int("ogreen", 25, "Outline green color compoment (0-255)")
	outlineBlue := flag.Int("oblue", 255, "Outline blue color compoment (0-255)")
	bgRed := flag.Int("bred", 255, "Background red color compoment (0-255)")
	bgGreen := flag.Int("bgreen", 255, "Background green color compoment (0-255)")
	bgBlue := flag.Int("bblue", 255, "Background blue color compoment (0-255)")
	of := flag.String("out", "out", "Output filename")

	flag.Parse()

	config := config.Config{
		Modulo: *mod,
		CanvasWidth: *cw, CanvasHeight: *ch, Radius: *r,
		LineWidth: *lw, OutlineWidth: *ow, 
		LineColor:[]int{*lineRed, *lineGreen, *lineBlue},
		OutlineColor: []int{*outlineRed, *outlineGreen, *outlineBlue},
		BackgroundColor: []int{*bgRed, *bgGreen, *bgBlue},
		OutputFile: *of, OutputType: "png",
	}

	return &config
}

func main() {
	config:= initArgs()
	draw.Draw(*config, period.Pisano(config.Modulo))
}
