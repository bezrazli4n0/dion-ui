package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/user32"
)

// Init инициализирует библиотеку
func Init() {
	initGraphics()
	appRunning = true
	dionWindows = make(map[*window]interface{})
}

var appRunning bool = false

// Exit выходит из приложения
func Exit() {
	user32.PostQuitMessage(0)
	appRunning = false
}

// Run запускает цикл обработки сообщений
func Run() int {
	msg := user32.MSG{}

	for appRunning {
		if user32.PeekMessage(&msg, 0, 0, 0, user32.PM_REMOVE) {
			user32.TranslateMessage(&msg)
			user32.DispatchMessage(&msg)
		} else {
			for w, _ := range dionWindows {
				w.render()
			}
		}
	}

	for w, _ := range dionWindows {
		w.Close()
	}

	freeGraphics()
	return int(msg.WParam)
}