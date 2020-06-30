package kernel32

import (
	"syscall"
)

var (
	kernel32DLL = syscall.NewLazyDLL("kernel32.dll")

	pGetModuleHandle = kernel32DLL.NewProc("GetModuleHandleW")
)

// GetModuleHandle возвращает дескриптор указанного модуля
func GetModuleHandle() HMODULE {
	ret, _, _ := pGetModuleHandle.Call(uintptr(0))
	return HMODULE(ret)
}