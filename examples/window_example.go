package main

import (
	"github.com/bezrazli4n0/dion-ui/pkg/app"
)

func main() {
	dion.Init()

	wnd, _ := dion.NewWindow("DionUI", 0, 0, 640, 480)
	wnd.AttachCallback(dion.OnClose, func() {
		wnd.Close()
	})

	dion.Run()
}