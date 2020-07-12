package main

import dion "github.com/bezrazli4n0/dion-ui/pkg/app"

func main() {
	dion.Init()

	wnd, _ := dion.NewWindow("Grid layout example", 0, 0, 400, 200, nil)
	wnd.AttachCallback(dion.OnClose, func() {
		dion.Exit()
	})

	btn := dion.NewButton("Hello - 1", 0.0, 0.0, 400.0, 200.0, 12.0, nil)
	btn1 := dion.NewButton("Hello - 2", 0.0, 0.0, 400.0, 200.0, 12.0, nil)
	btn2 := dion.NewButton("Hello - 3", 0.0, 0.0, 400.0, 200.0, 12.0, nil)

	internGrid := dion.NewGridLayout(0.0, 0.0, 2, 1)
	internGrid.AddWidget(btn1, 1, 1)
	internGrid.AddWidget(btn2, 2, 1)

	grid := dion.NewGridLayout(0.0, 0.0, 1, 2)
	grid.AddWidget(btn, 1, 1)
	grid.AddWidget(internGrid, 1, 2)

	wnd.SetWidget(grid)

	dion.Run()
}