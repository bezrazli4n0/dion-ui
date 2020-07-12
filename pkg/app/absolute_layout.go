package dion

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"

type AbsoluteLayout interface {
	Layout
	AddWidget(w widget)
}

type absoluteLayoutImpl struct {
	layoutImpl
}

func (l *absoluteLayoutImpl) draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight, parentX, parentY float32) {
	if l.widgets != nil && l.GetVisible() {
		for _, obj := range l.widgets {
			obj.w.draw(pRT, parentWidth, parentHeight, parentX, parentY)
		}
	}
}

func (l *absoluteLayoutImpl) AddWidget(w widget) {
	l.widgets = append(l.widgets, baseLayoutWidget{w: w})
}

func NewAbsoluteLayout() AbsoluteLayout {
	layout := &absoluteLayoutImpl{}
	layout.SetVisible(true)
	return layout
}