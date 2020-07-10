package d2d1

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/dwrite"
	"math"
	"syscall"
	"unsafe"
)

// ID2D1RenderTarget предоставляет объект, который обрабатывает команды рисования
type ID2D1RenderTarget struct {
	ID2D1Resource
}

// vtblID2D1RenderTarget виртуальная таблица для ID2D1RenderTarget
type vtblID2D1RenderTarget struct {
	vtblID2D1Resource
	CreateBitmap uintptr
	CreateBitmapFromWicBitmap uintptr
	CreateSharedBitmap uintptr
	CreateBitmapBrush uintptr
	CreateSolidColorBrush uintptr
	CreateGradientStopCollection uintptr
	CreateLinearGradientBrush uintptr
	CreateRadialGradientBrush uintptr
	CreateCompatibleRenderTarget uintptr
	CreateLayer uintptr
	CreateMesh uintptr
	DrawLine uintptr
	DrawRectangle uintptr
	FillRectangle uintptr
	DrawRoundedRectangle uintptr
	FillRoundedRectangle uintptr
	DrawEllipse uintptr
	FillEllipse uintptr
	DrawGeometry uintptr
	FillGeometry uintptr
	FillMesh uintptr
	FillOpacityMask uintptr
	DrawBitmap uintptr
	DrawText uintptr
	DrawTextLayout uintptr
	DrawGlyphRun uintptr
	SetTransform uintptr
	GetTransform uintptr
	SetAntialiasMode uintptr
	GetAntialiasMode uintptr
	SetTextAntialiasMode uintptr
	GetTextAntialiasMode uintptr
	SetTextRenderingParams uintptr
	GetTextRenderingParams uintptr
	SetTags uintptr
	GetTags uintptr
	PushLayer uintptr
	PopLayer uintptr
	Flush uintptr
	SaveDrawingState uintptr
	RestoreDrawingState uintptr
	PushAxisAlignedClip uintptr
	PopAxisAlignedClip uintptr
	Clear uintptr
	BeginDraw uintptr
	EndDraw uintptr
	GetPixelFormat uintptr
	SetDpi uintptr
	GetDpi uintptr
	GetSize uintptr
	GetPixelSize uintptr
	GetMaximumBitmapSize uintptr
	IsSupported uintptr
}

// vmt возвращает указатель на виртуальную таблицу
func (obj *ID2D1RenderTarget) vmt() *vtblID2D1RenderTarget {
	return (*vtblID2D1RenderTarget)(obj.Vtbl)
}

// BeginDraw https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-begindraw
func (obj *ID2D1RenderTarget) BeginDraw() {
	syscall.Syscall(obj.vmt().BeginDraw, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
}

// DrawText https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawtext(constwchar_uint32_idwritetextformat_constd2d1_rect_f__id2d1brush_d2d1_draw_text_options_dwrite_measuring_mode)
func (obj *ID2D1RenderTarget) DrawText(text string, textFormat *dwrite.IDWriteTextFormat, layoutRect RECT_F, defaultFillBrush *ID2D1Brush) {
	s, _ := syscall.UTF16FromString(text)

	syscall.Syscall9(
		obj.vmt().DrawText,
		8,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&s[0])),
		uintptr(len(s)),
		uintptr(unsafe.Pointer(textFormat)),
		uintptr(unsafe.Pointer(&layoutRect)),
		uintptr(unsafe.Pointer(defaultFillBrush)),
		uintptr(0),
		uintptr(0),
		0,
	)
}

// Clear https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-clear(constd2d1_color_f_)
func (obj *ID2D1RenderTarget) Clear(clearColor COLOR_F) {
	syscall.Syscall(obj.vmt().Clear, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&clearColor)), 0)
}

// FillRectangle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillrectangle(constd2d1_rect_f__id2d1brush)
func (obj *ID2D1RenderTarget) FillRectangle(rect RECT_F, brush *ID2D1Brush) {
	syscall.Syscall(obj.vmt().FillRectangle, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&rect)), uintptr(unsafe.Pointer(brush)))
}

// FillGeometry https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-fillgeometry
func (obj *ID2D1RenderTarget) FillGeometry(geometry *ID2D1Geometry, brush, opacityBrush *ID2D1Brush) {
	syscall.Syscall6(obj.vmt().FillGeometry, 4, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(geometry)), uintptr(unsafe.Pointer(brush)), uintptr(unsafe.Pointer(opacityBrush)), 0, 0)
}

// DrawRectangle https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawrectangle(constd2d1_rect_f__id2d1brush_float_id2d1strokestyle)
func (obj *ID2D1RenderTarget) DrawRectangle(rect RECT_F, brush *ID2D1Brush, strokeWidth float32) {
	syscall.Syscall6(obj.vmt().DrawRectangle, 5, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&rect)), uintptr(unsafe.Pointer(brush)), uintptr(math.Float32bits(strokeWidth)), uintptr(0), 0)
}

// DrawTextLayout https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-drawtextlayout
func (obj *ID2D1RenderTarget) DrawTextLayout(origin POINT_2F, textLayout *dwrite.IDWriteTextLayout, defaultFillBrush *ID2D1Brush) {
	syscall.Syscall6(obj.vmt().DrawTextLayout, 5, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&origin))), uintptr(unsafe.Pointer(textLayout)), uintptr(unsafe.Pointer(defaultFillBrush)), uintptr(0), 0)
}

// CreateSolidColorBrush https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-createsolidcolorbrush(constd2d1_color_f__constd2d1_brush_properties__id2d1solidcolorbrush)
func (obj *ID2D1RenderTarget) CreateSolidColorBrush(color COLOR_F, solidColorBrush **ID2D1SolidColorBrush) winapi.HRESULT {
	ret, _, _ := syscall.Syscall6(obj.vmt().CreateSolidColorBrush, 4, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&color)), uintptr(0), uintptr(unsafe.Pointer(solidColorBrush)), 0, 0)
	return winapi.HRESULT(ret)
}

// EndDraw https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1rendertarget-enddraw
func (obj *ID2D1RenderTarget) EndDraw() winapi.HRESULT {
	ret, _, _ := syscall.Syscall(obj.vmt().EndDraw, 3, uintptr(unsafe.Pointer(obj)), uintptr(0), uintptr(0))
	return winapi.HRESULT(ret)
}