package config

type Config struct {
	Modulo uint
	CanvasWidth, CanvasHeight int
	Radius float64
	LineWidth, OutlineWidth float64
	LineColor, OutlineColor, BackgroundColor []int
	OutputFile,OutputType string
}
