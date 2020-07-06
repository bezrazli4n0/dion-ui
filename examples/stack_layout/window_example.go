package main

import dion "github.com/bezrazli4n0/dion-ui/pkg/app"

func main() {
	dion.Init()

	wnd, _ := dion.NewWindow("Stack layout example", 0, 0, 400, 200)
	wnd.AttachCallback(dion.OnClose, func() {
		dion.Exit()
	})

	lbl := dion.NewLabel("Layout sample - 1", "Arial", 0.0, 0.0, 400, 200, 16.0, dion.Color{0, 0, 0, 255})
	lbl.SetTextAlignment(dion.LabelLeftH, dion.LabelCenterV)

	lbl1 := dion.NewLabel("Layout sample - 2", "Arial", 0.0, 0.0, 400, 200, 16.0, dion.Color{0, 0, 0, 255})
	lbl1.SetTextAlignment(dion.LabelCenterH, dion.LabelCenterV)

	lbl2 := dion.NewLabel("Layout sample - 3", "Arial", 0.0, 0.0, 400, 200, 16.0, dion.Color{0, 0, 0, 255})
	lbl2.SetTextAlignment(dion.LabelRightH, dion.LabelCenterV)

	stack := dion.NewStackLayout(0.0, 0.0, dion.StackV)
	stack.AddWidget(lbl)
	stack.AddWidget(lbl1)
	stack.AddWidget(lbl2)

	wnd.SetWidget(stack)

	dion.Run()
}
