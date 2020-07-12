package main

import dion "github.com/bezrazli4n0/dion-ui/pkg/app"

func main() {
	dion.Init()

	wndRoot, _ := dion.NewWindow("Root window", 0.0, 0.0, 400, 200, nil)
	wndRoot.AttachCallback(dion.OnClose, func() {
		dion.Exit()
	})

	btn := dion.NewButton("Show dialog", 10.0, 10.0, 100.0, 30.0, 12.0, func() {
		wndRoot.Enable(false)

		wndChild, _ := dion.NewWindow("Child window", 0, 0, 300, 150, wndRoot)
		wndChild.AttachCallback(dion.OnClose, func() {
			wndRoot.Enable(true)
		})

		btn1 := dion.NewButton("Close", 0.0, 0.0, 300.0, 150.0, 12.0, func() {
			wndChild.Close()
		})

		lbl := dion.NewLabel("Hello, world!", "Arial", 0.0, 0.0, 300.0, 150.0, 22.0, dion.Color{37, 37, 37, 255})
		lbl.SetTextAlignment(dion.LabelCenterH, dion.LabelCenterV)

		grid := dion.NewGridLayout(0.0, 0.0, 2, 1)
		grid.AddWidget(lbl, 1, 1)
		grid.AddWidget(btn1, 2, 1)

		wndChild.SetWidget(grid)
	})

	wndRoot.SetWidget(btn)

	dion.Run()
}