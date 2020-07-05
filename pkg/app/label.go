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
)

const (
	LabelLeftV LabelVerticalAlign = iota
	LabelRightV
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
	lbl.visible = true

	lbl.SetText(text)
	lbl.SetTextColor(textColor)
	lbl.SetFontName(fontName)
	lbl.SetFontSize(fontSize)
	lbl.SetSize(width, height)
	lbl.SetPos(x, y)

	return lbl
}

type labelImpl struct {
	widgetImpl

	textColor d2d1.COLOR_F
	text, fontName string
	fontSize float32
	hAlign LabelHorizontalAlign
	vAlign LabelVerticalAlign
}

func (obj *labelImpl) draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight float32) {
	if obj.visible {
		obj.calculateLayout(parentWidth, parentHeight)

		pTextFormat := &dwrite.IDWriteTextFormat{}
		getDWriteFactory().CreateTextFormat(obj.fontName, obj.fontSize, &pTextFormat)

		pTextFormat.SetTextAlignment(dwrite.TEXT_ALIGNMENT(obj.hAlign))
		pTextFormat.SetParagraphAlignment(dwrite.PARAGRAPH_ALIGNMENT(obj.vAlign))

		pBrush := &d2d1.ID2D1SolidColorBrush{}
		pRT.CreateSolidColorBrush(obj.textColor, &pBrush)

		pRT.DrawText(obj.text, pTextFormat, d2d1.RECT_F{pixelToDipX(obj.x), pixelToDipY(obj.y), pixelToDipX(obj.x + obj.width), pixelToDipY(obj.y + obj.height)}, (*d2d1.ID2D1Brush)(unsafe.Pointer(pBrush)))

		pBrush.Release()
		pTextFormat.Release()
	}
}

// SetFontName устанавливает шрифт
func (obj *labelImpl) SetFontName(fontName string) {
	obj.fontName = fontName
}

// SetFontSize устанавливает размер шрифта
func (obj *labelImpl) SetFontSize(fontSize float32) {
	obj.fontSize = fontSize
}

// SetText устанавливает текст
func (obj *labelImpl) SetText(text string) {
	obj.text = text
}

// SetTextAlignment выравнивает текст
func (obj *labelImpl) SetTextAlignment(hAlign LabelHorizontalAlign, vAlign LabelVerticalAlign) {
	obj.hAlign = hAlign
	obj.vAlign = vAlign
}

// SetTextColor устанавливает цвет текста
func (obj *labelImpl) SetTextColor(textColor Color) {
	obj.textColor = rgbaColorToColorF(textColor)
}