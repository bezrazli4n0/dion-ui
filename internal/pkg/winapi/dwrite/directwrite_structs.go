package dwrite

type TEXT_METRICS struct {
	Left float32
	Top float32
	Width float32
	WidthIncludingTrailingWhitespace float32
	Height float32
	LayoutWidth float32
	LayoutHeight float32
	MaxBidiReorderingDepth uint32
	LineCount uint32
}