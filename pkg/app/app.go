package dion

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/user32"

// Init инициализирует библиотеку
func Init() {
}

// Run запускает цикл обработки сообщений
func Run() int {
	uMsg := user32.MSG{}

	for user32.GetMessage(&uMsg, 0, 0, 0) {
		user32.TranslateMessage(&uMsg)
		user32.DispatchMessage(&uMsg)
	}

	return int(uMsg.WParam)
}