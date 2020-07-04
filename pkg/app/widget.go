package dion

type widget interface {
	graphicsObject

	SetVisible(visible bool)
}

type widgetImpl struct {
	graphicsObjectImpl

	visible bool
}

func (obj *widgetImpl) SetVisible(visible bool) {
	obj.visible = visible
}