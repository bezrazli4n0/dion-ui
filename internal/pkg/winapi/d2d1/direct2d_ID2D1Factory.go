package d2d1

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/combase"
	"syscall"
	"unsafe"
)

// ID2D1Factory содержит Direct2D ресурсы
type ID2D1Factory struct {
	vtbl *iD2D1FactoryVtbl
}

// iD2D1FactoryVtbl виртуальная таблица для ID2D1Factory
type iD2D1FactoryVtbl struct {
	combase.IUnknown

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

func (obj *ID2D1Factory) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

func (obj *ID2D1Factory) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// CreateHwndRenderTarget создаёт render target для обычного окна
func (obj *ID2D1Factory) CreateHwndRenderTarget(renderTargetProperties RENDER_TARGET_PROPERTIES, hwndRenderTargetProperties HWND_RENDER_TARGET_PROPERTIES, ppID2D1HwndRenderTarget **ID2D1HwndRenderTarget) winapi.HRESULT {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.CreateHwndRenderTarget,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(&renderTargetProperties)),
		uintptr(unsafe.Pointer(&hwndRenderTargetProperties)),
		uintptr(unsafe.Pointer(ppID2D1HwndRenderTarget)),
		0,
		0,
	)
	return winapi.HRESULT(ret)
}

// CreateFactory инициализирует Direct2D
func CreateFactory(factoryType FACTORY_TYPE, ppD2D1Factory **ID2D1Factory) winapi.HRESULT {
	ret, _, _ := pD2D1CreateFactory.Call(uintptr(factoryType), uintptr(unsafe.Pointer(&IID_ID2D1Factory)), uintptr(0), uintptr(unsafe.Pointer(ppD2D1Factory)))
	return winapi.HRESULT(ret)
}