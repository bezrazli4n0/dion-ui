package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/kernel32"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/user32"
	"math"
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

// updateAllWindows обновляет все окна движка
func updateAllWindows() {
	for w, _ := range dionWindows {
		w.update()
	}
}

// renderAllWindows перерисовывает все окна движка
func renderAllWindows() {
	for w, _ := range dionWindows {
		w.render()
	}
}

// Run запускает цикл обработки сообщений
func Run() int {
	msg := user32.MSG{}

	var nextUpdate int64 = 0
	const fps = 60
	const fpsInterval = 1000 / fps
	for appRunning {
		if user32.PeekMessage(&msg, 0, 0, 0, user32.PM_REMOVE) {
			user32.TranslateMessage(&msg)
			user32.DispatchMessage(&msg)
		} else {
			renderAllWindows()

			nowTime := time.Now()
			nextTime := nextUpdate

			var lWait int64 = 0

			if nowTime.Unix() < int64(nextTime) {
				lWait = int64(math.Min(float64(fpsInterval), float64(nextTime - nowTime.Unix())))
			}

			if lWait <= 1 {
				nextUpdate = nowTime.Unix() + fpsInterval
				updateAllWindows()
			} else {
				if user32.MsgWaitForMultipleObjects(0, nil, 0, winapi.DWORD(lWait), user32.QS_ALLEVENTS) == kernel32.WAIT_TIMEOUT {
					nextUpdate = time.Now().Unix() + fpsInterval
					updateAllWindows()
				}
			}
		}

	}

	for w, _ := range dionWindows {
		w.Close()
	}

	freeGraphics()
	return int(msg.WParam)
}