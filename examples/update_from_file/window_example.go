package main

import (
	"github.com/bezrazli4n0/dion-ui/pkg/app"
	"time"
)

func main() {
	dion.Init()

	wndState, _ := dion.NewWindowEngine(nil)
	wndState.LoadUIFromFileWithInterval("window.xml", 1 * time.Second)

	dion.Run()
}