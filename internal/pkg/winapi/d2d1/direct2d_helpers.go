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
