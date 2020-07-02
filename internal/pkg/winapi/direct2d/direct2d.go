package d2d1

import (
	"syscall"
)

var (
	d2d1DLL = syscall.NewLazyDLL("d2d1.dll")

	pD2D1CreateFactory = d2d1DLL.NewProc("D2D1CreateFactory")
)