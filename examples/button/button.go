package main

import (
	dion "github.com/bezrazli4n0/dion-ui/pkg/app"
	"log"
)

func main() {
	dion.Init()

	wnd, _ := dion.NewWindow("Button example", 0, 0, 400, 200)
	wnd.AttachCallback(dion.OnClose, func() {
		dion.Exit()
	})

	btn := dion.NewButton("Click me!", 10.0, 10.0, 100.0, 30.0, 12.0,  func() {
		log.Println("clicked!")
	})

	stack := dion.NewStackLayout(0.0, 0.0, dion.StackV)
	stack.AddWidget(btn)

	wnd.SetWidget(stack)

	dion.Run()
}
