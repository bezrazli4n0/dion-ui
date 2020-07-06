package dwrite

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"syscall"
	"unsafe"
)

type IDWriteTextLayout struct {
	IDWriteTextFormat
}

// vtblIDWriteTextLayout виртуальная таблица для IDWriteTextLayout
type vtblIDWriteTextLayout struct {
	vtblIDWriteTextFormat
	SetMaxWidth uintptr
	SetMaxHeight uintptr
	SetFontCollection uintptr
	SetFontFamilyName uintptr
	SetFontWeight uintptr
	SetFontStyle uintptr
	SetFontStretch uintptr
	SetFontSize uintptr
	SetUnderline uintptr
	SetStrikethrough uintptr
	SetDrawingEffect uintptr
	SetInlineObject uintptr
	SetTypography uintptr
	SetLocaleName uintptr
	GetMaxWidth uintptr
	GetMaxHeight uintptr
	GetFontCollection uintptr
	GetFontFamilyNameLength uintptr
	GetFontFamilyName uintptr
	GetFontWeight uintptr
	GetFontStyle uintptr
	GetFontStretch uintptr
	GetFontSize uintptr
	GetUnderline uintptr
	GetStrikethrough uintptr
	GetDrawingEffect uintptr
	GetInlineObject uintptr
	GetTypography uintptr
	GetLocaleNameLength uintptr
	GetLocaleName uintptr
	Draw uintptr
	GetLineMetrics uintptr
	GetMetrics uintptr
	GetOverhangMetrics uintptr
	GetClusterMetrics uintptr
	DetermineMinWidth uintptr
	HitTestPoint uintptr
	HitTestTextPosition uintptr
	HitTestTextRange uintptr
}

// vmt возвращает указатель на виртуальную таблицу
func (obj *IDWriteTextLayout) vmt() *vtblIDWriteTextLayout {
	return (*vtblIDWriteTextLayout)(obj.Vtbl)
}

// GetMetrics https://docs.microsoft.com/en-us/windows/win32/api/dwrite/nf-dwrite-idwritetextlayout-getmetrics
func (obj *IDWriteTextLayout) GetMetrics(textMetrics *TEXT_METRICS) winapi.HRESULT {
	ret, _, _ := syscall.Syscall(obj.vmt().GetMetrics, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(textMetrics)), 0)
	return winapi.HRESULT(ret)
}