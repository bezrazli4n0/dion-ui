package dwrite

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/com"
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