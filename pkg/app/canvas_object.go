package dion

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"

type CanvasObject interface {
	SetSize(width, height float32)
	SetPos(x, y float32)
	setColor(color d2d1.COLOR_F)
	SetColorRGB(r, g, b, a byte)

	draw(pRT *d2d1.ID2D1RenderTarget)
}

type canvasObjectImpl struct {
	width, height, x, y float32
	color d2d1.COLOR_F
}

func (obj *canvasObjectImpl) SetSize(width, height float32) {
	obj.width = width
	obj.height = height
}

func (obj *canvasObjectImpl) SetPos(x, y float32) {
	obj.x = x
	obj.y = y
}

func (obj *canvasObjectImpl) setColor(color d2d1.COLOR_F) {
	obj.color = color
}

func (obj *canvasObjectImpl) SetColorRGB(r, g, b, a byte) {
	obj.color = d2d1.COLOR_F{float32(r) / 255.0, float32(g) / 255.0, float32(b) / 255.0, float32(a) / 255.0}
}