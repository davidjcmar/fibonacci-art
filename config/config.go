package config

// enum for config.Config RGB255 selection
const (
	Red = iota
	Green = iota
	Blue = iota
)

type Config struct {
	Modulo uint
	CanvasWidth, CanvasHeight int
	Radius float64
	LineWidth, OutlineWidth float64
	LineColor, OutlineColor, BackgroundColor []int
	OutputFile,OutputType string
}
