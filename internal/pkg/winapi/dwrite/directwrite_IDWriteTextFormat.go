package dwrite

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/com"
	"syscall"
	"unsafe"
)

// IDWriteTextFormat
type IDWriteTextFormat struct {
	com.IUnknown
}

// vtblIDWriteTextFormat виртуальная таблица для IDWriteTextFormat
type vtblIDWriteTextFormat struct {
	com.VtblIUnknown
	SetTextAlignment uintptr
	SetParagraphAlignment uintptr
	SetWordWrapping uintptr
	SetReadingDirection uintptr
	SetFlowDirection uintptr
	SetIncrementalTabStop uintptr
	SetTrimming uintptr
	SetLineSpacing uintptr
	GetTextAlignment uintptr
	GetParagraphAlignment uintptr
	GetWordWrapping uintptr
	GetReadingDirection uintptr
	GetFlowDirection uintptr
	GetIncrementalTabStop uintptr
	GetTrimming uintptr
	GetLineSpacing uintptr
	GetFontCollection uintptr
	GetFontFamilyNameLength uintptr
	GetFontFamilyName uintptr
	GetFontWeight uintptr
	GetFontStyle uintptr
	GetFontStretch uintptr
	GetFontSize uintptr
	GetLocaleNameLength uintptr
	GetLocaleName uintptr
}

// vmt возвращает указатель на виртуальную таблицу
func (obj *IDWriteTextFormat) vmt() *vtblIDWriteTextFormat {
	return (*vtblIDWriteTextFormat)(obj.Vtbl)
}

// SetParagraphAlignment https://docs.microsoft.com/en-us/windows/win32/api/dwrite/nf-dwrite-idwritetextformat-setparagraphalignment
func (obj *IDWriteTextFormat) SetParagraphAlignment(paragraphAlignment PARAGRAPH_ALIGNMENT) winapi.HRESULT {
	ret, _, _ := syscall.Syscall(obj.vmt().SetParagraphAlignment, 2, uintptr(unsafe.Pointer(obj)), uintptr(paragraphAlignment), 0)
	return winapi.HRESULT(ret)
}

// SetTextAlignment https://docs.microsoft.com/en-us/windows/win32/api/dwrite/nf-dwrite-idwritetextformat-settextalignment
func (obj *IDWriteTextFormat) SetTextAlignment(textAlignment TEXT_ALIGNMENT) winapi.HRESULT {
	ret, _, _ := syscall.Syscall(obj.vmt().SetTextAlignment, 2, uintptr(unsafe.Pointer(obj)), uintptr(textAlignment), 0)
	return winapi.HRESULT(ret)
}