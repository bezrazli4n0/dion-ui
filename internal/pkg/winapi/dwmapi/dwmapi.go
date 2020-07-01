package dwmapi

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/user32"
	"syscall"
)

var (
	dwmapiDLL = syscall.NewLazyDLL("dwmapi.dll")

	pDwmGetWindowAttribute = dwmapiDLL.NewProc("DwmGetWindowAttribute")
)

// DwmGetWindowAttribute получает DWM атрибут окна
func DwmGetWindowAttribute(hWnd user32.HWND, dwAttribute winapi.DWORD, pvAttribute winapi.PVOID, cbAttribute winapi.DWORD) DWMAPI {
	ret, _, _ := pDwmGetWindowAttribute.Call(uintptr(hWnd), uintptr(dwAttribute), uintptr(pvAttribute), uintptr(cbAttribute))
	return DWMAPI(ret)
}