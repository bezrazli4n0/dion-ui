package app

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/kernel32"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/user32"
	"log"
	"syscall"
	"unsafe"
	"fmt"
)

// Window определяет внешний интерфейс взаимодействия
type Window interface {
	SetTitle(title string)
	SetPos(x, y int)
	SetSize(width, height int)
	GetHandle() user32.HWND
}

// window содержит реализацию
type window struct {
	hWnd user32.HWND
	title, class string
}

// SetTitle устанавливает заголовок окна
func (w *window) SetTitle(title string) {
	log.Println("Not implemented yet")
}

// SetPos устанавливает позицию окна
func (w *window) SetPos(x, y int) {
	log.Println("Not implemented yet")
}

// SetSize устанавливает размеры окна
func (w *window) SetSize(width, height int) {
	log.Println("Not implemented yet")
}

// GetHandle возвращает handle окна
func (w *window) GetHandle() user32.HWND {
	return w.hWnd
}

func NewWindow(title string, x, y, width, height int) Window {

	wndProc := func(hWnd user32.HWND, uMsg uint32, wParam user32.WPARAM, lParam user32.LPARAM) user32.LRESULT {
		switch uMsg {
		case user32.WM_CLOSE:
			user32.PostQuitMessage(0)
			return 0
		}
		return user32.DefWindowProc(hWnd, uMsg, wParam, lParam)
	}

	wc := user32.WNDCLASSEXW{
		LpszClassName: syscall.StringToUTF16Ptr(fmt.Sprintf("%s_dionUI", title)),
		LpfnWndProc: syscall.NewCallback(wndProc),
		HCursor: user32.LoadCursor(user32.IDC_ARROW),
		HInstance: user32.HINSTANCE(kernel32.GetModuleHandle()),
		HbrBackground: user32.COLOR_WINDOW + 1,
	}
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	user32.RegisterClassEx(&wc)

	hWnd := user32.CreateWindowEx(0, syscall.StringToUTF16Ptr(fmt.Sprintf("%s_dionUI", title)), syscall.StringToUTF16Ptr(title), user32.WS_OVERLAPPEDWINDOW, int32(x), int32(y), int32(width), int32(height), 0, 0, wc.HInstance)

	user32.UpdateWindow(hWnd)
	user32.ShowWindow(hWnd, user32.SW_SHOW)

	return &window{hWnd: hWnd, title: title, class: fmt.Sprintf("%s_dionUI", title)}
}
