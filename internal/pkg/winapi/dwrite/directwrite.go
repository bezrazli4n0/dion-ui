package dwrite

import "syscall"

var (
	d2d1DLL = syscall.NewLazyDLL("dwrite.dll")

	pDWriteCreateFactory = d2d1DLL.NewProc("DWriteCreateFactory")
)