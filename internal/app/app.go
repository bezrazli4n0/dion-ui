package app

import "github.com/bezrazli4n0/dion-ui/internal/pkg/user32"

func Run() {
	uMsg := user32.MSG{}

	for user32.GetMessage(&uMsg, 0, 0, 0) {
		user32.TranslateMessage(&uMsg)
		user32.DispatchMessage(&uMsg)
	}
}