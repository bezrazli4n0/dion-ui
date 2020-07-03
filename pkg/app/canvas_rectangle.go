package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
	"unsafe"
)

func NewRectangle(x, y, width, height float32, filled bool, color Color) CanvasObject {
	if filled {
		fillRect := &fillRectangle{}
		fillRect.SetPos(x, y)
		fillRect.SetSize(width, height)
		fillRect.SetColorRGB(color.R, color.G, color.B, color.A)
		return fillRect
	}

	rect := &drawRectangle{}
	rect.SetPos(x, y)
	rect.SetSize(width, height)
	rect.SetColorRGB(color.R, color.G, color.B, color.A)
	return rect
}

type fillRectangle struct {
	canvasObjectImpl
}

type drawRectangle struct {
	canvasObjectImpl
}

func (obj *fillRectangle) draw(pRT *d2d1.ID2D1RenderTarget) {
	pBrush := &d2d1.ID2D1SolidColorBrush{}
	pRT.CreateSolidColorBrush(obj.color, &pBrush)
	pRT.FillRectangle(d2d1.RECT_F{pixelToDipX(obj.x), pixelToDipY(obj.y), pixelToDipX(obj.x + obj.width), pixelToDipY(obj.y + obj.height)}, (*d2d1.ID2D1Brush)(unsafe.Pointer(pBrush)))
	pBrush.Release()
}

func (obj *drawRectangle) draw(pRT *d2d1.ID2D1RenderTarget) {
	pBrush := &d2d1.ID2D1SolidColorBrush{}
	pRT.CreateSolidColorBrush(obj.color, &pBrush)
	pRT.DrawRectangle(d2d1.RECT_F{pixelToDipX(obj.x + 0.5), pixelToDipY(obj.y + 0.5), pixelToDipX(obj.x + obj.width + 0.5), pixelToDipY(obj.y + obj.height + 0.5)}, (*d2d1.ID2D1Brush)(unsafe.Pointer(pBrush)), 1.0)
	pBrush.Release()
}
