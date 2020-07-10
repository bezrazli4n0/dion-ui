package dion

type mouseEventType int32

const (
	onMouseMove mouseEventType = iota

	onRButtonUp
	onRButtonDown

	onLButtonUp
	onLButtonDown

	onMouseEnter
	onMouseLeave
)

type widget interface {
	graphicsObject

	SetVisible(visible bool)
	GetVisible() bool
	getMinBounds() (float32, float32)
	Dispose()

	needHandleMouse(x, y int, eventType mouseEventType)
	isHandleMouse() bool
	isMouseInside() bool

	onMouseMove(x, y int)

	onRButtonUp(x, y int)
	onRButtonDown(x, y int)

	onLButtonUp(x, y int)
	onLButtonDown(x, y int)

	onMouseEnter(x, y int)
	onMouseLeave(x, y int)
}

type widgetImpl struct {
	graphicsObjectImpl

	visible, handleMouse, mouseInside bool
	minWidth, minHeight float32
}

func (obj *widgetImpl) needHandleMouse(x, y int, eventType mouseEventType) {
	obj.handleMouse = false
}

func (obj *widgetImpl) isHandleMouse() bool {
	return obj.handleMouse
}

func (obj *widgetImpl) isMouseInside() bool {
	return obj.mouseInside
}

func (obj *widgetImpl) onMouseMove(x, y int) {
}

func (obj *widgetImpl) onRButtonUp(x, y int) {
}

func (obj *widgetImpl) onRButtonDown(x, y int) {
}

func (obj *widgetImpl) onLButtonUp(x, y int) {
}

func (obj *widgetImpl) onLButtonDown(x, y int) {
}

func (obj *widgetImpl) onMouseEnter(x, y int) {

}

func (obj *widgetImpl) onMouseLeave(x, y int) {

}

func (obj *widgetImpl) getMinBounds() (float32, float32) {
	return obj.minWidth, obj.minHeight
}

func (obj *widgetImpl) GetVisible() bool {
	return obj.visible
}

func (obj *widgetImpl) Dispose() {
}

func (obj *widgetImpl) SetVisible(visible bool) {
	obj.visible = visible
}