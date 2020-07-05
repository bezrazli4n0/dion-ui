package main

import (
	"github.com/bezrazli4n0/dion-ui/pkg/app"
)

func main() {
	dion.Init()

	wnd, _ := dion.NewWindow("Label example", 0, 0, 400, 200)
	wnd.AttachCallback(dion.OnClose, func() {
		dion.Exit()
	})

	lbl := dion.NewLabel("Привет", "Arial", 0, 0, 400, 200, 16, dion.Color{37, 37, 37, 255})
	lbl.SetTextAlignment(dion.LabelCenterH, dion.LabelCenterV)
	wnd.SetWidget(lbl)

	dion.Run()
}