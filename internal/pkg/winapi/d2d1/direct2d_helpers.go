package d2d1

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/user32"

func RenderTargetProperties() RENDER_TARGET_PROPERTIES {
	return RENDER_TARGET_PROPERTIES {
		Type: 0,
		PixelFormat: PIXEL_FORMAT{
			AlphaMode: 0,
			Format: 0,
		},
		DpiX: 0,
		DpiY: 0,
		Usage: 0,
		MinLevel: 0,
	}
}

func IdentityMatrix3x2() MATRIX_3X2_F {
	mtx := MATRIX_3X2_F{}
	mtx.U.S1.M11 = 1.0
	mtx.U.S1.M22 = 1.0

	mtx.U.S2.F11 = 1.0
	mtx.U.S2.F22 = 1.0

	mtx.U.M[0][0] = 1.0
	mtx.U.M[1][1] = 1.0

	return mtx
}

func HwndRenderTargetProperties(hWnd user32.HWND, width, height int) HWND_RENDER_TARGET_PROPERTIES {
	return HWND_RENDER_TARGET_PROPERTIES{
		Hwnd: hWnd,
		PixelSize: SIZE_U{
			Width: uint32(width),
			Height: uint32(height),
		},
		PresentOptions: 0,
	}
}
