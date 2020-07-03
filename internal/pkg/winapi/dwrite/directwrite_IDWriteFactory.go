package dwrite

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/com"
	"math"
	"syscall"
	"unsafe"
)

// IDWriteFactory используется для создания DirectWrite объектов
type IDWriteFactory struct {
	com.IUnknown
}

// vtblIDWriteFactory виртуальная таблица для IDWriteFactory
type vtblIDWriteFactory struct {
	com.VtblIUnknown
	GetSystemFontCollection uintptr
	CreateCustomFontCollection uintptr
	RegisterFontCollectionLoader uintptr
	UnregisterFontCollectionLoader uintptr
	CreateFontFileReference uintptr
	CreateCustomFontFileReference uintptr
	CreateFontFace uintptr
	CreateRenderingParams uintptr
	CreateMonitorRenderingParams uintptr
	CreateCustomRenderingParams uintptr
	RegisterFontFileLoader uintptr
	UnregisterFontFileLoader uintptr
	CreateTextFormat uintptr
	CreateTypography uintptr
	GetGdiInterop uintptr
	CreateTextLayout uintptr
	CreateGdiCompatibleTextLayout uintptr
	CreateEllipsisTrimmingSign uintptr
	CreateTextAnalyzer uintptr
	CreateNumberSubstitution uintptr
	CreateGlyphRunAnalysis uintptr
}

// vmt возвращает указатель на виртуальную таблицу
func (obj *IDWriteFactory) vmt() *vtblIDWriteFactory {
	return (*vtblIDWriteFactory)(obj.Vtbl)
}

// CreateTextFormat создаёт формат текста для рендера
func (obj *IDWriteFactory) CreateTextFormat(fontFamilyName string, fontSize float32, ppIDWriteTextFormat **IDWriteTextFormat) winapi.HRESULT {
	_text, _ := syscall.UTF16FromString(fontFamilyName)
	_localText, _ := syscall.UTF16FromString("")

	ret, _, _ := syscall.Syscall9(
		obj.vmt().CreateTextFormat,
		9,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&_text[0])),
		uintptr(0),
		uintptr(400),
		uintptr(0),
		uintptr(5),
		uintptr(math.Float32bits(fontSize)),
		uintptr(unsafe.Pointer(&_localText[0])),
		uintptr(unsafe.Pointer(ppIDWriteTextFormat)),
	)
	return winapi.HRESULT(ret)
}

// CreateTextLayout создаёт объект который содержит всю информацию о тексте
func (obj *IDWriteFactory) CreateTextLayout(text string, pTextFormat *IDWriteTextFormat, maxWidth, maxHeight float32, ppIDWriteTextLayout **IDWriteTextLayout) winapi.HRESULT {
	_text, _ := syscall.UTF16FromString(text)

	ret, _, _ := syscall.Syscall9(
		obj.vmt().CreateTextLayout,
		7,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&_text[0])),
		uintptr(len(_text)),
		uintptr(unsafe.Pointer(pTextFormat)),
		uintptr(math.Float32bits(maxWidth)),
		uintptr(math.Float32bits(maxHeight)),
		uintptr(unsafe.Pointer(ppIDWriteTextLayout)),
		0,
		0,
	)
	return winapi.HRESULT(ret)
}

// CreateFactory инициализирует DirectWrite
func CreateFactory(factoryType FACTORY_TYPE, ppIDWriteFactory **IDWriteFactory) winapi.HRESULT {
	ret, _, _ := pDWriteCreateFactory.Call(uintptr(factoryType), uintptr(unsafe.Pointer(&IID_IDWriteFactory)), uintptr(unsafe.Pointer(ppIDWriteFactory)))
	return winapi.HRESULT(ret)
}
