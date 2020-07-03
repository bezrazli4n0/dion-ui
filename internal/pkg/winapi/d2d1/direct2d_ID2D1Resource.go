package d2d1

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/com"

// ID2D1Resource предоставляет Direct2D ресурсы для рисования
type ID2D1Resource struct {
	com.IUnknown
}

// vtblID2D1Resource виртуальная таблица для ID2D1Resource
type vtblID2D1Resource struct {
	com.VtblIUnknown
	GetFactory uintptr
}