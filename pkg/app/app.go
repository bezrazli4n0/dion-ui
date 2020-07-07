package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/user32"
	"time"
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

	go func() {
		for {
			for w, _ := range dionWindows {
				w.invalidate()
			}
			time.Sleep(60 / time.Second)
		}
	}()

	for appRunning {
		if user32.GetMessage(&msg, 0, 0, 0) {
			user32.TranslateMessage(&msg)
			user32.DispatchMessage(&msg)
		}
	}

	for w, _ := range dionWindows {
		w.Close()
	}

	freeGraphics()
	return int(msg.WParam)
}