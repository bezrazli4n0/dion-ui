package main

import (
	"github.com/bezrazli4n0/dion-ui/pkg/app"
)

func main() {
	dion.Init()

	wnd, _ := dion.NewWindow("Label example", 0, 0, 400, 200, nil)
	wnd.AttachCallback(dion.OnClose, func() {
		dion.Exit()
	})

	lbl := dion.NewLabel("Привет, мир! - こんにちは世界! - Hello, world!", "Arial", 0.0, 0.0, 400.0, 200.0, 16.0, dion.Color{0, 0, 0, 255})
	lbl.SetTextAlignment(dion.LabelCenterH, dion.LabelCenterV)

	wnd.SetWidget(lbl)

	dion.Run()
}