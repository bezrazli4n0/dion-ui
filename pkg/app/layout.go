package dion

type Layout interface {
	widget

	AddWidget(wdgt widget)
}

type layoutImpl struct {
	widgetImpl
	widgets []widget
}

func (l *layoutImpl) needHandleMouse(x, y int, eventType mouseEventType) {
	l.handleMouse = true

	for _, obj := range l.widgets {
		if obj.GetVisible() {
			obj.needHandleMouse(x, y, eventType)
		}
	}
}

func (l *layoutImpl) onLButtonUp(x, y int) {
	for _, obj := range l.widgets {
		if obj.GetVisible() && obj.isHandleMouse() {
			obj.onLButtonUp(x, y)
		}
	}
}

func (l *layoutImpl) onLButtonDown(x, y int) {
	for _, obj := range l.widgets {
		if obj.GetVisible() && obj.isHandleMouse() {
			obj.onLButtonDown(x, y)
		}
	}
}

func (l *layoutImpl) onRButtonUp(x, y int) {
	for _, obj := range l.widgets {
		if obj.GetVisible() && obj.isHandleMouse() {
			obj.onRButtonUp(x, y)
		}
	}
}

func (l *layoutImpl) onRButtonDown(x, y int) {
	for _, obj := range l.widgets {
		if obj.GetVisible() && obj.isHandleMouse() {
			obj.onRButtonDown(x, y)
		}
	}
}

func (l *layoutImpl) onMouseEnter(x, y int) {
	l.mouseInside = true
	for _, obj := range l.widgets {
		if obj.GetVisible() && obj.isHandleMouse() && !obj.isMouseInside() {
			obj.onMouseEnter(x, y)
		}
	}
}

func (l *layoutImpl) onMouseLeave(x, y int) {
	l.mouseInside = false
	for _, obj := range l.widgets {
		if obj.GetVisible() && obj.isHandleMouse() && obj.isMouseInside() {
			obj.onMouseLeave(x, y)
		}
	}
}

func (l *layoutImpl) onMouseMove(x, y int) {
	for _, obj := range l.widgets {
		if obj.GetVisible() && obj.isHandleMouse() {
			obj.onMouseMove(x, y)
		}
	}
}
