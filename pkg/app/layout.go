package dion

type Layout interface {
	widget
}

type baseLayoutWidget struct {
	w widget
	winfo interface{}
}

type layoutImpl struct {
	widgetImpl
	widgets []baseLayoutWidget
}

func (l *layoutImpl) update() {
	for _, obj := range l.widgets {
		if obj.w.GetVisible() {
			obj.w.update()
		}
	}
}

func (l *layoutImpl) needHandleMouse(x, y int, eventType mouseEventType) {
	l.handleMouse = true

	for _, obj := range l.widgets {
		if obj.w.GetVisible() {
			obj.w.needHandleMouse(x, y, eventType)
		}
	}
}

func (l *layoutImpl) onLButtonUp(x, y int) {
	for _, obj := range l.widgets {
		if obj.w.GetVisible() && obj.w.isHandleMouse() {
			obj.w.onLButtonUp(x, y)
		}
	}
}

func (l *layoutImpl) getMinBounds() (float32, float32) {
	l.minWidth = 0
	l.minHeight = 0
	if l.widgets != nil {
		for _, obj := range l.widgets {
			minWidth, minHeight := obj.w.getMinBounds()
			l.minWidth += minWidth
			l.minHeight += minHeight
		}
	}
	return l.minWidth, l.minHeight
}

func (l *layoutImpl) onLButtonDown(x, y int) {
	for _, obj := range l.widgets {
		if obj.w.GetVisible() && obj.w.isHandleMouse() {
			obj.w.onLButtonDown(x, y)
		}
	}
}

func (l *layoutImpl) onRButtonUp(x, y int) {
	for _, obj := range l.widgets {
		if obj.w.GetVisible() && obj.w.isHandleMouse() {
			obj.w.onRButtonUp(x, y)
		}
	}
}

func (l *layoutImpl) onRButtonDown(x, y int) {
	for _, obj := range l.widgets {
		if obj.w.GetVisible() && obj.w.isHandleMouse() {
			obj.w.onRButtonDown(x, y)
		}
	}
}

func (l *layoutImpl) onMouseEnter(x, y int) {
	l.mouseInside = true
	for _, obj := range l.widgets {
		if obj.w.GetVisible() && obj.w.isHandleMouse() && !obj.w.isMouseInside() {
			obj.w.onMouseEnter(x, y)
		}
	}
}

func (l *layoutImpl) onMouseLeave(x, y int) {
	l.mouseInside = false

	for _, obj := range l.widgets {
		if (obj.w.GetVisible() && obj.w.isHandleMouse() && obj.w.isMouseInside()) || (obj.w.GetVisible() && obj.w.isHandleMouse()) {
			obj.w.onMouseLeave(x, y)
		}
	}
}

func (l *layoutImpl) onMouseMove(x, y int) {
	for _, obj := range l.widgets {
		if obj.w.GetVisible() && obj.w.isHandleMouse() {
			obj.w.onMouseMove(x, y)
		}
	}
}
