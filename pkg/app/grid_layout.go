package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
	"log"
)

// TODO: improve grid paddings

// TODO: add column and row span option

// TODO: add x and y offsets
type GridLayout interface {
	Layout
	AddWidget(w widget, row, col uint32)
	SetPadding(x, y float32)
}

func NewGridLayout(x, y float32, rows, cols uint32) GridLayout {
	layout := &gridLayoutImpl{}
	layout.rows = rows
	layout.cols = cols

	layout.SetPadding(2.0, 2.0)
	layout.SetPos(x, y)
	layout.SetVisible(true)
	return layout
}

type gridLayoutImpl struct {
	layoutImpl
	rows, cols uint32
	paddingX, paddingY float32
}

type gridWidget struct {
	row, col uint32
}

func (l *gridLayoutImpl) SetPadding(x, y float32) {
	l.paddingX = x
	l.paddingY = y
}

func (l *gridLayoutImpl) draw(pRT *d2d1.ID2D1RenderTarget, parentWidth, parentHeight, parentX, parentY float32) {
	l.width = parentWidth
	l.height = parentHeight
	l.x = parentX
	l.y = parentY

	if l.widgets != nil {
		cellWidth := uint32(parentWidth) / l.cols
		cellHeight := uint32(parentHeight) / l.rows

		for _, obj := range l.widgets {
			if obj.w.GetVisible() {
				cellX := float32(cellWidth) * float32(obj.winfo.(gridWidget).col - 1) + parentX
				cellY := float32(cellHeight) * float32(obj.winfo.(gridWidget).row - 1) + parentY

				obj.w.SetSize(float32(cellWidth) - l.paddingX * 2, float32(cellHeight) - l.paddingY * 2)
				obj.w.SetPos(cellX + l.paddingX, cellY + l.paddingY)

				obj.w.draw(pRT, float32(cellWidth), float32(cellHeight), cellX, cellY)
			}
		}
	}
}

func (l *gridLayoutImpl) getMinBounds() (float32, float32) {
	return l.width, l.height
}

func (l *gridLayoutImpl) AddWidget(w widget, row, col uint32) {
	if row > l.rows || col > l.cols {
		log.Fatal("row or col overflow grid: rows - ", l.rows, " cols - ", l.cols, ", row - ", row, " col - ", col)
	}

	l.widgets = append(l.widgets, baseLayoutWidget{w: w, winfo: gridWidget{row, col}})
}
