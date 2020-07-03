package main

import (
	"github.com/bezrazli4n0/dion-ui/pkg/app"
)

func main() {
	dion.Init()

	wndState, _ := dion.NewWindowEngine(nil)
	wndState.LoadUIFromFile("window.xml")

	dion.Run()
}