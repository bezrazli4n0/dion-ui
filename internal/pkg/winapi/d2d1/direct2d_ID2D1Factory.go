package d2d1

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/com"
	"syscall"
	"unsafe"
)

// ID2D1Factory содержит Direct2D ресурсы
type ID2D1Factory struct {
	com.IUnknown
}

// vtblID2D1Factory виртуальная таблица для ID2D1Factory
type vtblID2D1Factory struct {
	com.VtblIUnknown
	ReloadSystemMetrics uintptr
	GetDesktopDpi uintptr
	CreateRectangleGeometry uintptr
	CreateRoundedRectangleGeometry uintptr
	CreateEllipseGeometry uintptr
	CreateGeometryGroup uintptr
	CreateTransformedGeometry uintptr
	CreatePathGeometry uintptr
	CreateStrokeStyle uintptr
	CreateDrawingStateBlock uintptr
	CreateWicBitmapRenderTarget uintptr
	CreateHwndRenderTarget uintptr
	CreateDxgiSurfaceRenderTarget uintptr
	CreateDCRenderTarget uintptr
}

// vmt возвращает указатель на виртуальную таблицу
func (f *ID2D1Factory) vmt() *vtblID2D1Factory {
	return (*vtblID2D1Factory)(f.Vtbl)
}

// CreateHwndRenderTarget создаёт render target для обычного окна
func (f *ID2D1Factory) CreateHwndRenderTarget(renderTargetProperties RENDER_TARGET_PROPERTIES, hwndRenderTargetProperties HWND_RENDER_TARGET_PROPERTIES, ppID2D1HwndRenderTarget **ID2D1HwndRenderTarget) winapi.HRESULT {
	ret, _, _ := syscall.Syscall6(
		f.vmt().CreateHwndRenderTarget,
		4,
		uintptr(unsafe.Pointer(f)),
		uintptr(unsafe.Pointer(&renderTargetProperties)),
		uintptr(unsafe.Pointer(&hwndRenderTargetProperties)),
		uintptr(unsafe.Pointer(ppID2D1HwndRenderTarget)),
		0,
		0,
	)
	return winapi.HRESULT(ret)
}

// GetDesktopDpi возвращает dpi монитора
func (f *ID2D1Factory) GetDesktopDpi() (dpiX, dpiY float32) {
	syscall.Syscall(f.vmt().GetDesktopDpi, 3, uintptr(unsafe.Pointer(f)), uintptr(unsafe.Pointer(&dpiX)), uintptr(unsafe.Pointer(&dpiY)))
	return
}

// CreateFactory инициализирует Direct2D
func CreateFactory(factoryType FACTORY_TYPE, ppD2D1Factory **ID2D1Factory) winapi.HRESULT {
	ret, _, _ := pD2D1CreateFactory.Call(uintptr(factoryType), uintptr(unsafe.Pointer(&IID_ID2D1Factory)), uintptr(0), uintptr(unsafe.Pointer(ppD2D1Factory)))
	return winapi.HRESULT(ret)
}