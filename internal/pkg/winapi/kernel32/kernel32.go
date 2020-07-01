package kernel32

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"syscall"
)

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")

	pGetModuleHandle = kernel32DLL.NewProc("GetModuleHandleW")
	pGetLastError = kernel32DLL.NewProc("GetLastError")
	pSetLastError = kernel32DLL.NewProc("SetLastError")
)

// GetModuleHandle возвращает дескриптор указанного модуля
func GetModuleHandle() HMODULE {
	ret, _, _ := pGetModuleHandle.Call(uintptr(0))
	return HMODULE(ret)
}

// GetLastError возвращает идентификатор ошибки
func GetLastError() winapi.DWORD {
	ret, _, _ := pGetLastError.Call()
	return winapi.DWORD(ret)
}

// SetLastError устанавливает код последней ошибки для потока
func SetLastError(dwErrorCode winapi.DWORD) {
	pSetLastError.Call(uintptr(dwErrorCode))
}