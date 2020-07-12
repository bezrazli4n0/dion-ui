package main

import (
	dion "github.com/bezrazli4n0/dion-ui/pkg/app"
	"log"
)

func main() {
	dion.Init()

	wnd, _ := dion.NewWindow("Absolute layout example", 0, 0, 300, 200, nil)
	wnd.AttachCallback(dion.OnClose, func() {
		dion.Exit()
	})

	lbl := dion.NewLabel("Hello, world!", "Arial", 10.0, 10.0, 100.0, 15.0, 12.0, dion.Color{37, 37, 37, 255})

	btn := dion.NewButton("Click Me!", 10.0, 30.0, 100.0, 30.0, 12.0, func() {
		log.Println("Clicked!")
	})

	btn1 := dion.NewButton("Click Me - 1!", 10.0, 65.0, 100.0, 30.0, 12.0, func() {
		log.Println("Clicked - 1!")
	})

	absoluteLayout := dion.NewAbsoluteLayout()
	absoluteLayout.AddWidget(lbl)
	absoluteLayout.AddWidget(btn)
	absoluteLayout.AddWidget(btn1)

	wnd.SetWidget(absoluteLayout)
	dion.Run()
}
