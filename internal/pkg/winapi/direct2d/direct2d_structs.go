package d2d1

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/user32"

type SIZE_U struct {
	Width uint32
	Height uint32
}

type PIXEL_FORMAT struct {
	Format int32
	AlphaMode int32
}

type RENDER_TARGET_PROPERTIES struct {
	Type int32
	PixelFormat PIXEL_FORMAT
	DpiX float32
	DpiY float32
	Usage int32
	MinLevel int32
}

type COLOR_F struct {
	R float32
	G float32
	B float32
	A float32
}

type HWND_RENDER_TARGET_PROPERTIES struct {
	Hwnd user32.HWND
	PixelSize SIZE_U
	PresentOptions int32
}