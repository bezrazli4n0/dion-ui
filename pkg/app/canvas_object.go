package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
)

type CanvasObject interface {
	graphicsObject

	SetColorRGBA(color Color)
	setColor(color d2d1.COLOR_F)
}

type canvasObjectImpl struct {
	graphicsObjectImpl
	color d2d1.COLOR_F
}

func (obj *canvasObjectImpl) setColor(color d2d1.COLOR_F) {
	obj.color = color
}

func (obj *canvasObjectImpl) SetColorRGBA(color Color) {
	obj.color = rgbaColorToColorF(color)
}