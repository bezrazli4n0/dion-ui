package d2d1

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/combase"
	"syscall"
	"unsafe"
)

// ID2D1HwndRenderTarget рендерер для рисования в окне
type ID2D1HwndRenderTarget struct {
	vtbl *iD2D1HwndRenderTargetVtbl
}

// iD2D1HwndRenderTargetVtbl виртуальная таблица для ID2D1HwndRenderTarget
type iD2D1HwndRenderTargetVtbl struct {
	combase.IUnknown

	ID2D1Resource

	ID2D1RenderTarget

	CheckWindowState uintptr
	Resize uintptr
	GetHwnd uintptr
}

func (obj *ID2D1HwndRenderTarget) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

func (obj *ID2D1HwndRenderTarget) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// EndDraw операции рисования
func (obj *ID2D1HwndRenderTarget) EndDraw() winapi.HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.EndDraw,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(0),
		uintptr(0),
	)
	return winapi.HRESULT(ret)
}

// Clear очищает область рисования указанным цветом
func (obj *ID2D1HwndRenderTarget) Clear(clearColor COLOR_F) {
	syscall.Syscall(
		obj.vtbl.Clear,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&clearColor)),
		0,
	)
}

// Resize изменяет размер render target
func (obj *ID2D1HwndRenderTarget) Resize(pixelSize SIZE_U) winapi.HRESULT {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Resize,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&pixelSize)),
		0,
	)
	return winapi.HRESULT(ret)
}

// BeginDraw инициализирует начало рисования на render target
func (obj *ID2D1HwndRenderTarget) BeginDraw() {
	syscall.Syscall(
		obj.vtbl.BeginDraw,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
}