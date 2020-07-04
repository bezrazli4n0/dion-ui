package dion

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"

func pixelToDipX(x float32) float32 {
	return (x * dpiX) / 96.0
}

func pixelToDipY(y float32) float32 {
	return (y * dpiY) / 96.0
}

func dipToPixelX(x float32) float32 {
	return (x * 96.0) / dpiX
}

func dipToPixelY(y float32) float32 {
	return (y * 96.0) / dpiY
}

func rgbaColorToColorF(color Color) d2d1.COLOR_F {
	return d2d1.COLOR_F{float32(color.R) / 255.0, float32(color.G) / 255.0, float32(color.B) / 255.0, float32(color.A) / 255.0}
}