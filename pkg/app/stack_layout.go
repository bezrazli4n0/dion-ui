package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
)

type StackLayoutOrientation byte

const (
	StackH StackLayoutOrientation = iota
	StackV
)

// TODO: implement layout padding
func NewStackLayout(x, y float32, orientation StackLayoutOrientation) Layout {
	layout := &stackLayoutImpl{}
	layout.SetPos(x, y)
	layout.SetVisible(true)
	layout.orientation = orientation
	layout.marginX = 2.0
	layout.marginY = 2.0
	return layout
}

type stackLayoutImpl struct {
	layoutImpl
	orientation StackLayoutOrientation
	marginX, marginY float32
}

func (l *stackLayoutImpl) getMinBounds() (float32, float32) {
	l.minWidth = 0
	l.minHeight = 0
	if l.widgets != nil {
		for _, obj := range l.widgets {
			minWidth, minHeight := obj.getMinBounds()
			l.minWidth += minWidth
			l.minHeight += minHeight
		}
	}
	return l.minWidth, l.minHeight
}

func (l *stackLayoutImpl) AddWidget(wdgt widget) {
	l.widgets = append(l.widgets, wdgt)
	l.width, l.height = l.getMinBounds()
}

func (l *stackLayoutImpl) draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight, parentX, parentY float32) {
	if l.widgets != nil {
		deltaX := l.x + parentX + l.marginX
		deltaY := l.y + parentY + l.marginY

		switch l.orientation {
		case StackV:
			l.width = parentWidth
			for _, obj := range l.widgets {
				if obj.GetVisible() {
					_, minHeight := obj.getMinBounds()

					obj.SetSize(parentWidth, minHeight)
					obj.SetPos(deltaX - l.marginX, deltaY)

					obj.draw(pRT, l.width, l.height, l.x, l.y)

					deltaY += minHeight + l.marginY
				}
			}
			break
		case StackH:
			l.height = parentHeight
			for _, obj := range l.widgets {
				if obj.GetVisible() {
					minWidth, _ := obj.getMinBounds()

					obj.SetSize(minWidth, parentHeight)
					obj.SetPos(deltaX, deltaY - l.marginY)
					obj.draw(pRT, l.width, l.height, l.x, l.y)

					deltaX += minWidth + l.marginX
				}
			}
			break
		}
	}
}