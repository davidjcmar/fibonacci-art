package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	config "github.com/davidjcmar/fibonacci-art/config"
	period "github.com/davidjcmar/fibonacci-art/period"
)

func help() {
	fmt.Println(os.Args[0], "[mod]")
}

func initArgs() *config.Config {
	cw := flag.Int("width", 250, "Canvas width")
	ch := flag.Int("height", 250, "Canvas height")
	r := flag.Float64("radius", 120, "Circle radius")
	lw := flag.Float64("line", 2, "Line width")
	ow := flag.Float64("outline", 2, "Outline width")
	lineRed := flag.Int("red", 0, "Line red color component (0-255)")
	lineGreen := flag.Int("green", 0, "Line green component (0-255)")
	lineBlue := flag.Int("blue", 0, "Line blue component (0-255)")
	outlineRed := flag.Int("ored", 0, "Outline red color compoment (0-255)")
	outlineGreen := flag.Int("ogreen", 0, "Outline green color compoment (0-255)"
	outlineBlue := flag.Int("oblue", 0, "Outline blue color compoment (0-255)"
	bgRed := flag.Int("bred", 255, "Background red color compoment (0-255)")
	bgGreen := flag.Int("bgreen", 255, "Background green color compoment (0-255)")
	bgBlue := flag.Int("bblue", 255, "Background blue color compoment (0-255)")
	of := flag.String("o", "out.png", "Output filename")

	config := config.Config{
		CanvasWidth: *cw, CanvasHeight: *ch,, Radius: *r,
		LineWidth: *lw, OutlineWidth: *ow, 
		LineRGBInt:[]int{*lineRed, *lineGreen, *lineBlue},
		OutlineRGBInt: []int{*outlineRed, *outlineGreen, *outlineBlue},
		BackgroundRGBInt: []int{*bgRed, *bgGreen, *bgBlue},
		OutputFile: of, OutputType: "png"
	}

	return &config
}

func main() {
	config:= initArgs()

	/*
	if len(os.Args) != 2 {
		help()
		os.Exit(-1)
	} else {
		u32, err := strconv.ParseUint(os.Args[1], 10, 32)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		p, _ := period.Pisano(uint(u32))
		fmt.Println(p)
	}
	*/
}
