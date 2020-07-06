package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
	"unsafe"
)

func NewRectangle(x, y, width, height float32, filled bool, color Color, strokeWidth float32) CanvasObject {
	if filled {
		fillRect := &fillRectangle{}
		fillRect.SetPos(x, y)
		fillRect.SetSize(width, height)
		fillRect.SetColorRGBA(color)
		return fillRect
	}

	rect := &drawRectangle{}
	rect.SetPos(x, y)
	rect.SetSize(width, height)
	rect.SetColorRGBA(color)
	if strokeWidth != 0.0 {
		rect.SetStrokeWidth(strokeWidth)
	} else {
		rect.SetStrokeWidth(1.0)
	}
	return rect
}

type fillRectangle struct {
	canvasObjectImpl
}

type drawRectangle struct {
	canvasObjectImpl
	strokeWidth float32
}

func (obj *drawRectangle) SetStrokeWidth(strokeWidth float32) {
	obj.strokeWidth = strokeWidth
}

func (obj *fillRectangle) draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight, parentX, parentY float32) {
	obj.calculateLayout(parentWidth, parentHeight)

	pBrush := &d2d1.ID2D1SolidColorBrush{}
	pRT.CreateSolidColorBrush(obj.color, &pBrush)
	pRT.FillRectangle(d2d1.RECT_F{pixelToDipX(obj.x), pixelToDipY(obj.y), pixelToDipX(obj.x + obj.width), pixelToDipY(obj.y + obj.height)}, (*d2d1.ID2D1Brush)(unsafe.Pointer(pBrush)))
	pBrush.Release()
}

func (obj *drawRectangle) draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight, parentX, parentY float32) {
	obj.calculateLayout(parentWidth, parentHeight)

	pBrush := &d2d1.ID2D1SolidColorBrush{}
	pRT.CreateSolidColorBrush(obj.color, &pBrush)
	pRT.DrawRectangle(d2d1.RECT_F{pixelToDipX(obj.x + 0.5), pixelToDipY(obj.y + 0.5), pixelToDipX(obj.x + obj.width + 0.5), pixelToDipY(obj.y + obj.height + 0.5)}, (*d2d1.ID2D1Brush)(unsafe.Pointer(pBrush)), obj.strokeWidth)
	pBrush.Release()
}
