package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
	"math"
)

type graphicsObject interface {
	calculateLayout(parentWidth, parentHeight float32)

	getPercentSize() (widthPercent, heightPercent float32)
	GetSize() (width, height float32)
	GetPos() (x, y float32)

	setPercentSize(widthPercent, heightPercent float32)
	SetSize(width, height float32)
	SetPos(x, y float32)

	draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight float32)
}

type graphicsObjectImpl struct {
	width, height, x, y, widthPercent, heightPercent float32
}

func (obj *graphicsObjectImpl) calculateLayout(parentWidth, parentHeight float32) {
	widthPercent, heightPercent := obj.getPercentSize()
	width, height := obj.GetSize()

	if widthPercent > 0 {
		width = (parentWidth / 100.0) * widthPercent
	}

	if heightPercent > 0 {
		height = (parentHeight / 100.0) * heightPercent
	}

	obj.SetSize(float32(math.Round(float64(width))), float32(math.Round(float64(height))))
}

func (obj *graphicsObjectImpl) GetSize() (width, height float32) {
	return obj.width, obj.height
}

func (obj *graphicsObjectImpl) getPercentSize() (widthPercent, heightPercent float32) {
	return obj.widthPercent, obj.heightPercent
}

func (obj *graphicsObjectImpl) setPercentSize(widthPercent, heightPercent float32) {
	obj.widthPercent = widthPercent
	obj.heightPercent = heightPercent
}

func (obj *graphicsObjectImpl) SetSize(width, height float32) {
	obj.width = width
	obj.height = height
}

func (obj *graphicsObjectImpl) SetPos(x, y float32) {
	obj.x = x
	obj.y = y
}

func (obj *graphicsObjectImpl) GetPos() (x, y float32) {
	return obj.x, obj.y
}