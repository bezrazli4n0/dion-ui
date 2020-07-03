package dion

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