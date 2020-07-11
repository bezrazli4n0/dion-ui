package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
	"unsafe"
)

type Button interface {
	widget
}

type buttonImpl struct {
	widgetImpl
	lblText Label

	roundRect *d2d1.ID2D1RoundedRectangleGeometry

	backgroundColor d2d1.COLOR_F
	backgroundColorHover d2d1.COLOR_F
	backgroundColorClick d2d1.COLOR_F

	onClickCallback interface{}

	isMouseDown bool

	text string
	x, y, width, height, fontSize float32
	textColor Color
}

// NewButton создает новую кнопку
func NewButton(text string, x, y, width, height, fontSize float32, onClick interface{}) Button {
	btn := &buttonImpl{}

	btn.x = x
	btn.y = y

	btn.width = width
	btn.height = height

	btn.backgroundColor = rgbaToColorF(0, 134, 224, 255)
	btn.backgroundColorHover = rgbaToColorF(0, 113, 189, 255)
	btn.backgroundColorClick = rgbaToColorF(0, 92, 153, 255)

	btn.onClickCallback = onClick
	btn.text = text
	btn.textColor = Color{255, 255, 255, 255}
	btn.fontSize = fontSize

	btn.recreateInternalResource()

	btn.SetVisible(true)

	return btn
}

// recreateInternalResource пересоздаёт внутренние ресурсы объекта
func (btn *buttonImpl) recreateInternalResource() {
	btn.Dispose()

	if pixelToDipX(btn.width) > 0 && pixelToDipY(btn.height) > 0 {
		rect := d2d1.ROUNDED_RECT{
			Rect:    d2d1.RECT_F{pixelToDipX(btn.x), pixelToDipY(btn.y), pixelToDipX(btn.x + btn.width), pixelToDipY(btn.y + btn.height)},
			RadiusX: pixelToDipX(2.0),
			RadiusY: pixelToDipY(2.0),
		}

		getD2D1Factory().CreateRoundedRectangleGeometry(rect, &btn.roundRect)

		btn.lblText = NewLabel(btn.text, "Arial", btn.x, btn.y, btn.width, btn.height, btn.fontSize, btn.textColor)
		btn.lblText.SetTextAlignment(LabelCenterH, LabelCenterV)
		minWidth, minHeight := btn.lblText.getMinBounds()
		btn.minWidth = minWidth + pixelToDipX(8.0)
		btn.minHeight = minHeight + pixelToDipY(10.0)
	}
}

// Dispose высвобождает память
func (btn *buttonImpl) Dispose() {
	if btn.lblText != nil {
		btn.lblText.Dispose()
		btn.lblText = nil
	}

	if btn.roundRect != nil {
		btn.roundRect.Release()
		btn.roundRect = nil
	}
}

// SetPos устанавливает позицию объекта
func (btn *buttonImpl) SetPos(x, y float32) {
	btn.x = x
	btn.y = y
	btn.recreateInternalResource()
}

// SetSize устанавливает ширину и высоту объекта
func (btn *buttonImpl) SetSize(width, height float32) {
	btn.width = width
	btn.height = height
	btn.recreateInternalResource()
}

// needHandleMouse проверяет, нужно ли обрабатывать события мыши
func (btn *buttonImpl) needHandleMouse(x, y int, eventType mouseEventType) {
	worldTransform := d2d1.IdentityMatrix3x2()
	_, contains := btn.roundRect.FillContainsPoint(d2d1.POINT_2F{float32(x), float32(y)}, &worldTransform, 1.0)

	switch eventType {
	case onMouseEnter:
		btn.handleMouse = contains
		break

	case onMouseLeave:
		btn.handleMouse = !contains
		break

	case onLButtonDown:
		btn.handleMouse = contains
		break

	case onLButtonUp:
		btn.handleMouse = contains
		break

	default:
		btn.handleMouse = false
	}
}

// onLButtonUp вызывается когда левая кнопка мыши отпущена
func (btn *buttonImpl) onLButtonUp(x, y int) {
	btn.isMouseDown = false

	if btn.onClickCallback != nil {
		btn.onClickCallback.(func())()
	}
}

// onLButtonDown вызывается когда левая кнопка мыши была нажата
func (btn *buttonImpl) onLButtonDown(x, y int) {
	btn.isMouseDown = true
}

// onMouseLeave срабатывает когда мышка покинула зону кнопки
func (btn *buttonImpl) onMouseLeave(x, y int) {
	btn.mouseInside = false
}

// onMouseEnter срабатывает когда мышка вошла в зону кнопки
func (btn *buttonImpl) onMouseEnter(x, y int) {
	btn.mouseInside = true
}

// draw рисует виджет
func (btn *buttonImpl) draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight, parentX, parentY float32) {
	if pixelToDipX(btn.width) > 0 && pixelToDipY(btn.height) > 0 {
		pBrush := &d2d1.ID2D1SolidColorBrush{}

		if btn.isMouseDown {
			pRT.CreateSolidColorBrush(btn.backgroundColorClick, &pBrush)
		} else if btn.mouseInside {
			pRT.CreateSolidColorBrush(btn.backgroundColorHover, &pBrush)
		} else {
			pRT.CreateSolidColorBrush(btn.backgroundColor, &pBrush)
		}

		pRT.FillGeometry((*d2d1.ID2D1Geometry)(unsafe.Pointer(btn.roundRect)), (*d2d1.ID2D1Brush)(unsafe.Pointer(pBrush)), nil)

		btn.lblText.draw(pRT, btn.width, btn.height, btn.x, btn.y)

		pBrush.Release()
	}
}