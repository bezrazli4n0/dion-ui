package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
	"math"
)

type StackLayoutOrientation byte

type StackLayout interface {
	Layout
	AddWidget(w widget)
	SetPadding(x, y float32)
}

const (
	StackH StackLayoutOrientation = iota
	StackV
)

func NewStackLayout(x, y float32, orientation StackLayoutOrientation) StackLayout {
	layout := &stackLayoutImpl{}
	layout.SetPos(x, y)
	layout.SetVisible(true)
	layout.orientation = orientation
	layout.paddingX = 2.0
	layout.paddingY = 2.0
	return layout
}

type stackLayoutImpl struct {
	layoutImpl
	orientation StackLayoutOrientation
	paddingX, paddingY float32
}

func (l *stackLayoutImpl) AddWidget(w widget) {
	l.widgets = append(l.widgets, baseLayoutWidget{w: w})
	l.width, l.height = l.getMinBounds()
}

func (l *stackLayoutImpl) SetPadding(x, y float32) {
	l.paddingX = x
	l.paddingY = y
}

func (l *stackLayoutImpl) draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight, parentX, parentY float32) {
	if l.widgets != nil && l.GetVisible() {
		deltaX := float32(math.Ceil(float64(l.x + parentX + l.paddingX)))
		deltaY := float32(math.Ceil(float64(l.y + parentY + l.paddingY)))

		switch l.orientation {
		case StackV:
			l.width = parentWidth
			for _, obj := range l.widgets {
				if obj.w.GetVisible() {
					_, minHeight := obj.w.getMinBounds()

					obj.w.SetSize(float32(math.Ceil(float64(parentWidth - l.paddingX * 2))), float32(math.Ceil(float64(minHeight))))
					obj.w.SetPos(deltaX, deltaY)

					obj.w.draw(pRT, float32(math.Ceil(float64(l.width))), float32(math.Ceil(float64(l.height))), float32(math.Ceil(float64(l.x))), float32(math.Ceil(float64(l.y))))

					deltaY += float32(math.Ceil(float64(minHeight + l.paddingY)))
				}
			}
			break
		case StackH:
			l.height = parentHeight
			for _, obj := range l.widgets {
				if obj.w.GetVisible() {
					minWidth, _ := obj.w.getMinBounds()

					obj.w.SetSize(float32(math.Ceil(float64(minWidth))), float32(math.Ceil(float64(parentHeight - l.paddingY * 2))))
					obj.w.SetPos(deltaX, deltaY)
					obj.w.draw(pRT, float32(math.Ceil(float64(l.width))), float32(math.Ceil(float64(l.height))), float32(math.Ceil(float64(l.x))), float32(math.Ceil(float64(l.y))))

					deltaX += float32(math.Ceil(float64(minWidth + l.paddingX)))
				}
			}
			break
		}
	}
}