package dion

type widget interface {
	graphicsObject

	SetVisible(visible bool)
	GetVisible() bool
	getMinBounds() (float32, float32)
	Dispose()
}

type widgetImpl struct {
	graphicsObjectImpl

	visible bool
	minWidth, minHeight float32
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