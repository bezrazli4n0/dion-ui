package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/dwrite"
	"unsafe"
)

type LabelHorizontalAlign int32
type LabelVerticalAlign int32

const (
	LabelLeftH LabelHorizontalAlign = iota
	LabelRightH
	LabelCenterH
	LabelJustifiedH
)

const (
	LabelTopV LabelVerticalAlign = iota
	LabelBottomV
	LabelCenterV
)

type Label interface {
	widget

	SetTextColor(textColor Color)
	SetText(text string)
	SetFontName(fontName string)
	SetFontSize(fontSize float32)
	SetTextAlignment(hAlign LabelHorizontalAlign, vAlign LabelVerticalAlign)
}

func NewLabel(text, fontName string, x, y, width, height, fontSize float32, textColor Color) Label {
	lbl := &labelImpl{}

	lbl.text = text
	lbl.fontName = fontName
	lbl.fontSize = fontSize

	lbl.hAlign = LabelLeftH
	lbl.vAlign = LabelTopV

	lbl.width = width
	lbl.height = height
	lbl.x = x
	lbl.y = y

	lbl.SetVisible(true)
	lbl.SetTextColor(textColor)

	lbl.recreateInternalResources()

	return lbl
}

type labelImpl struct {
	widgetImpl

	textColor d2d1.COLOR_F
	text, fontName string
	fontSize float32
	hAlign LabelHorizontalAlign
	vAlign LabelVerticalAlign

	pTextFormat *dwrite.IDWriteTextFormat
	pTextLayout *dwrite.IDWriteTextLayout
}

func (obj *labelImpl) recreateInternalResources() {
	obj.Dispose()

	getDWriteFactory().CreateTextFormat(obj.fontName, obj.fontSize, &obj.pTextFormat)

	obj.pTextFormat.SetTextAlignment(dwrite.TEXT_ALIGNMENT(obj.hAlign))
	obj.pTextFormat.SetParagraphAlignment(dwrite.PARAGRAPH_ALIGNMENT(obj.vAlign))

	getDWriteFactory().CreateTextLayout(obj.text, obj.pTextFormat, obj.width, obj.height, &obj.pTextLayout)

	textMetrics := dwrite.TEXT_METRICS{}
	obj.pTextLayout.GetMetrics(&textMetrics)
	obj.minWidth = textMetrics.Width
	obj.minHeight = textMetrics.Height
}

func (obj *labelImpl) draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight, parentX, parentY float32) {
	if obj.GetVisible() {
		obj.calculateLayout(parentWidth, parentHeight)

		pBrush := &d2d1.ID2D1SolidColorBrush{}
		pRT.CreateSolidColorBrush(obj.textColor, &pBrush)

		pRT.DrawTextLayout(d2d1.POINT_2F{obj.x, obj.y}, obj.pTextLayout, (*d2d1.ID2D1Brush)(unsafe.Pointer(pBrush)))

		pBrush.Release()
	}
}

// SetFontName устанавливает шрифт
func (obj *labelImpl) SetFontName(fontName string) {
	obj.fontName = fontName
	obj.recreateInternalResources()
}

// SetFontSize устанавливает размер шрифта
func (obj *labelImpl) SetFontSize(fontSize float32) {
	obj.fontSize = fontSize
	obj.recreateInternalResources()
}

// SetText устанавливает текст
func (obj *labelImpl) SetText(text string) {
	obj.text = text
	obj.recreateInternalResources()
}

// SetSize устанавливает ширину и высоту объекта
func (obj *labelImpl) SetSize(width, height float32) {
	obj.width = width
	obj.height = height
	obj.recreateInternalResources()
}

// SetPos устанавливает позицию объекта
func (obj *labelImpl) SetPos(x, y float32) {
	obj.x = x
	obj.y = y
	obj.recreateInternalResources()
}

// Dispose высвобождает память
func (obj *labelImpl) Dispose() {
	if obj.pTextLayout != nil {
		obj.pTextLayout.Release()
		obj.pTextLayout = nil
	}

	if obj.pTextFormat != nil {
		obj.pTextFormat.Release()
		obj.pTextFormat = nil
	}
}

// SetTextAlignment выравнивает текст
func (obj *labelImpl) SetTextAlignment(hAlign LabelHorizontalAlign, vAlign LabelVerticalAlign) {
	obj.hAlign = hAlign
	obj.vAlign = vAlign
	obj.recreateInternalResources()
}

// SetTextColor устанавливает цвет текста
func (obj *labelImpl) SetTextColor(textColor Color) {
	obj.textColor = rgbaColorToColorF(textColor)
}