package dion

type Layout interface {
	widget

	AddWidget(wdgt widget)
}

type layoutImpl struct {
	widgetImpl
	widgets []widget
}
