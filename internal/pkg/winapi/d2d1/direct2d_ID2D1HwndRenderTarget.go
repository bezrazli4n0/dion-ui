package d2d1

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"syscall"
	"unsafe"
)

// ID2D1HwndRenderTarget рендерер для рисования в окне
type ID2D1HwndRenderTarget struct {
	ID2D1RenderTarget
}

// vtblID2D1HwndRenderTarget виртуальная таблица для ID2D1HwndRenderTarget
type vtblID2D1HwndRenderTarget struct {
	vtblID2D1RenderTarget
	CheckWindowState uintptr
	Resize uintptr
	GetHwnd uintptr
}

// vmt возвращает указатель на виртуальную таблицу
func (obj *ID2D1HwndRenderTarget) vmt() *vtblID2D1HwndRenderTarget {
	return (*vtblID2D1HwndRenderTarget)(obj.Vtbl)
}

// Resize https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1hwndrendertarget-resize(constd2d1_size_u_)
func (obj *ID2D1HwndRenderTarget) Resize(pixelSize SIZE_U) winapi.HRESULT {
	ret, _, _ := syscall.Syscall(obj.vmt().Resize, 2, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(&pixelSize)), 0)
	return winapi.HRESULT(ret)
}