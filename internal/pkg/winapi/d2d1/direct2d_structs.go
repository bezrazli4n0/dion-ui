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

type RECT_F struct {
	Left float32
	Top float32
	Right float32
	Bottom float32
}

type POINT_2F struct {
	X float32
	Y float32
}

type ROUNDED_RECT struct {
	Rect RECT_F
	RadiusX float32
	RadiusY float32
}

type MATRIX_3X2_F struct {
	U struct {
		S1 struct {
			M11 float32
			M12 float32
			M21 float32
			M22 float32
			Dx float32
			Dy float32
		}

		S2 struct {
			F11 float32
			F12 float32
			F21 float32
			F22 float32
			F31 float32
			F32 float32
		}

		M [3][2]float32
	}
}