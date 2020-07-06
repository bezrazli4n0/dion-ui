package dion

import (
	"errors"
	"fmt"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/kernel32"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/user32"
	"syscall"
	"unsafe"
)

type WindowCallbackType int

const (
	OnLMouseButtonDown WindowCallbackType = iota
	OnLMouseButtonUp

	OnRMouseButtonDown
	OnRMouseButtonUp

	OnMouseMove

	OnClose

	OnResize
)

// Window определяет внешний интерфейс взаимодействия с окном
type Window interface {
	SetTitle(title string)
	SetPos(x, y int)
	SetSize(width, height int)

	GetHandle() user32.HWND
	GetTitle() string
	GetSize() (width, height int)
	GetPos() (x, y int)

	AttachCallback(callbackType WindowCallbackType, callback interface{})
	DetachCallback(callbackType WindowCallbackType)

	SetWidget(wdgt widget)

	SetBackgroundColor(r, g, b, a byte)
	SetCanvas(canvas *Canvas)

	Close()
	Hide()
	Show()
}

// window содержит реализацию окна
type window struct {
	hWnd         user32.HWND
	title, class string
	x, y, width, height int
	pRT *d2d1.ID2D1HwndRenderTarget
	backgroundColor d2d1.COLOR_F
	isClosed bool

	canvas *Canvas
	layoutWidget widget

	callbacks map[WindowCallbackType]interface{}
}

// dionWindows содержит все окна
var dionWindows map[*window]interface{}

// SetTitle устанавливает заголовок окна
func (w *window) SetTitle(title string) {
	if user32.SetWindowText(w.hWnd, title) {
		w.title = title
	}
}

// SetPos устанавливает позицию окна
func (w *window) SetPos(x, y int) {
	w.x = x
	w.y = y

	winRect := w.getCorrectedBox()
	x += int(winRect.Left)

	user32.SetWindowPos(w.hWnd, user32.HWND_TOP, x, y, 0, 0, user32.SWP_NOOWNERZORDER | user32.SWP_NOZORDER | user32.SWP_NOSIZE)
}

// SetSize устанавливает размеры окна
func (w *window) SetSize(width, height int) {
	w.width = width
	w.height = height

	winRect := w.getCorrectedBox()
	width = int(winRect.Right - winRect.Left)
	height = int(winRect.Bottom - winRect.Top)

	user32.SetWindowPos(w.hWnd, user32.HWND_TOP, 0, 0, width, height, user32.SWP_NOOWNERZORDER | user32.SWP_NOZORDER | user32.SWP_NOMOVE)
}

// GetHandle возвращает handle окна
func (w *window) GetHandle() user32.HWND {
	return w.hWnd
}

// AddWidget добавляет виджет в окно
func (w *window) SetWidget(wdgt widget) {
	w.layoutWidget = wdgt
}

// GetTitle возвращает заголовок окна
func (w *window) GetTitle() string {
	return w.title
}

// AttachCallback устанавливает функцию обратного вызова на события окна
func (w *window) AttachCallback(callbackType WindowCallbackType, callback interface{}) {
	w.callbacks[callbackType] = callback
}

// windowProc процедура передачи сообщений окнам
func windowProc(hWnd user32.HWND, uMsg uint32, wParam winapi.WPARAM, lParam winapi.LPARAM) winapi.LRESULT {
	var pWindow *window

	if uMsg == user32.WM_NCCREATE {
		dataPtr := (*user32.CREATESTRUCTW)(unsafe.Pointer(lParam)).LpCreateParams
		pWindow = (*window)(unsafe.Pointer(dataPtr))

		kernel32.SetLastError(0)
		if user32.SetWindowLongPtr(hWnd, user32.GWLP_USERDATA, winapi.LONG_PTR(unsafe.Pointer(pWindow))) > 0{
			if kernel32.GetLastError() != 0 {
				return 0
			}
		}
	} else {
		pWindow = (*window)(unsafe.Pointer(user32.GetWindowLongPtr(hWnd, user32.GWLP_USERDATA)))
	}

	if pWindow != nil {
		return pWindow.windowProc(hWnd, uMsg, wParam, lParam)
	}

	return user32.DefWindowProc(hWnd, uMsg, wParam, lParam)
}

// windowProc процедура обработки сообщений окна
func (w *window) windowProc(hWnd user32.HWND, uMsg uint32, wParam winapi.WPARAM, lParam winapi.LPARAM) winapi.LRESULT {
	switch uMsg {
	case user32.WM_DISPLAYCHANGE:
	case user32.WM_PAINT:
		ps := &user32.PAINTSTRUCT{}
		user32.BeginPaint(hWnd, ps)
		w.render()
		user32.EndPaint(hWnd, ps)
		return 0

	case user32.WM_LBUTTONUP:
		w.onLButtonUp(int(user32.Loword(int32(lParam))), int(user32.Hiword(int32(lParam))))
		return 0

	case user32.WM_LBUTTONDOWN:
		w.onLButtonDown(int(user32.Loword(int32(lParam))), int(user32.Hiword(int32(lParam))))
		return 0

	case user32.WM_RBUTTONUP:
		w.onRButtonUp(int(user32.Loword(int32(lParam))), int(user32.Hiword(int32(lParam))))
		return 0

	case user32.WM_RBUTTONDOWN:
		w.onRButtonDown(int(user32.Loword(int32(lParam))), int(user32.Hiword(int32(lParam))))
		return 0

	case user32.WM_MOUSEMOVE:
		w.onMouseMove(int(user32.Loword(int32(lParam))), int(user32.Hiword(int32(lParam))))
		return 0

	case user32.WM_SIZE:
		w.onResize(int(user32.Loword(int32(lParam))), int(user32.Hiword(int32(lParam))))
		return 0

	case user32.WM_DESTROY:
		if callback, ok := w.callbacks[OnClose]; ok {
			callback.(func())()
		}

		w.isClosed = true
		delete(dionWindows, w)

		return 0
	}
	return user32.DefWindowProc(hWnd, uMsg, wParam, lParam)
}

// onLButtonDown обрабатывает собитие нажатия левой кнопкой мыши
func (w *window) onLButtonDown(x, y int) {
	if callback, ok := w.callbacks[OnLMouseButtonDown]; ok {
		callback.(func(x, y int))(x, y)
	}
}

// onLButtonUp обрабатывает собитие отжатия левой кнопки мыши
func (w *window) onLButtonUp(x, y int) {
	if callback, ok := w.callbacks[OnLMouseButtonUp]; ok {
		callback.(func(x, y int))(x, y)
	}
}

// SetBackgroundColor устанавливает цвет окна
func (w *window) SetBackgroundColor(r, g, b, a byte) {
	w.backgroundColor = d2d1.COLOR_F{R: float32(r) / 255.0, G: float32(g) / 255.0, B: float32(b) / 255.0, A: float32(a) / 255.0}
}

// onRButtonDown обрабатывает собитие нажатия правой кнопкой мыши
func (w *window) onRButtonDown(x, y int) {
	if callback, ok := w.callbacks[OnRMouseButtonDown]; ok {
		callback.(func(x, y int))(x, y)
	}
}

// onRButtonUp обрабатывает собитие отжатия правой кнопки мыши
func (w *window) onRButtonUp(x, y int) {
	if callback, ok := w.callbacks[OnRMouseButtonUp]; ok {
		callback.(func(x, y int))(x, y)
	}
}

// onMouseMove обрабатывает движение мыши
func (w *window) onMouseMove(x, y int) {
	if callback, ok := w.callbacks[OnMouseMove]; ok {
		callback.(func(x, y int))(x, y)
	}
}

// onResize вызывается при изменении размеров окна
func (w *window) onResize(width, height int) {
	w.width = width
	w.height = height

	if w.pRT != nil {
		w.pRT.Resize(d2d1.SIZE_U{uint32(width), uint32(height)})
	}

	w.render()

	if callback, ok := w.callbacks[OnResize]; ok {
		callback.(func(width, height int))(width, height)
	}
}

// DetachCallback удаляет функцию обратного вызова с события
func (w *window) DetachCallback(callbackType WindowCallbackType) {
	_, ok := w.callbacks[callbackType]
	if ok {
		delete(w.callbacks, callbackType)
	}
}

// Close функция закрывает окно
func (w *window) Close() {
	user32.DestroyWindow(w.hWnd)
	user32.UnregisterClass(w.class, user32.HINSTANCE(kernel32.GetModuleHandle()))
	if w.pRT != nil {
		w.pRT.Release()
	}
}

// Hide скрывает окно
func (w *window) Hide() {
	user32.ShowWindow(w.hWnd, user32.SW_HIDE)
}

// Show показывает окно
func (w *window) Show() {
	user32.ShowWindow(w.hWnd, user32.SW_SHOW)
}

// GetPos возвращает позицию окна
func (w *window) GetPos() (int, int) {
	return w.x, w.y
}

// render отрисовывает окно
func (w *window) render() {
	w.pRT.BeginDraw()
	w.pRT.Clear(w.backgroundColor)

	// Draw canvas
	if w.canvas != nil {
		for _, obj := range w.canvas.Child {
			obj.draw((*d2d1.ID2D1RenderTarget)(unsafe.Pointer(w.pRT)), float32(w.width), float32(w.height), 0.0, 0.0)
		}
	}

	// Draw widgets
	if w.layoutWidget != nil {
		w.layoutWidget.draw((*d2d1.ID2D1RenderTarget)(unsafe.Pointer(w.pRT)), float32(w.width), float32(w.height), 0.0, 0.0)
	}

	w.pRT.EndDraw()
}

// GetSize возвращает размер окна
func (w *window) GetSize() (width, height int) {
	return w.width, w.height
}

// SetCanvas устанавливает Canvas для свободного рисования в окне
func (w *window) SetCanvas(canvas *Canvas) {
	w.canvas = canvas
}

// getCorrectedBox возвращает размеры действительного прямоугольника окна
func (w *window) getCorrectedBox() user32.RECT {
	winRect := user32.RECT{Right: int32(w.width), Bottom: int32(w.height)}
	user32.AdjustWindowRect(&winRect, user32.WS_OVERLAPPEDWINDOW, false)
	return winRect
}

// NewWindow возвращает новый экземпляр окна
func NewWindow(title string, x, y, width, height int) (Window, error) {
	wc := user32.WNDCLASSEXW{
		LpszClassName: syscall.StringToUTF16Ptr(fmt.Sprintf("%s_dionUI", title)),
		LpfnWndProc:   syscall.NewCallback(windowProc),
		HCursor:       user32.LoadCursor(user32.IDC_ARROW),
		HInstance:     user32.HINSTANCE(kernel32.GetModuleHandle()),
		HbrBackground: user32.COLOR_WINDOW + 1,
	}
	wc.CbSize = uint32(unsafe.Sizeof(wc))
	user32.RegisterClassEx(&wc)

	wnd := &window{title: title, class: fmt.Sprintf("%s_dionUI", title)}
	wnd.callbacks = make(map[WindowCallbackType]interface{})
	wnd.x = x
	wnd.y = y
	wnd.width = width
	wnd.height = height
	wnd.backgroundColor = d2d1.COLOR_F{1.0, 1.0, 1.0, 1.0}
	wnd.isClosed = false

	winRect := wnd.getCorrectedBox()
	x += int(winRect.Left)
	width = int(winRect.Right - winRect.Left)
	height = int(winRect.Bottom - winRect.Top)

	hWnd := user32.CreateWindowEx(0, wnd.class, title, user32.WS_OVERLAPPEDWINDOW, int32(x), int32(y), int32(width), int32(height), 0, 0, wc.HInstance, winapi.LPVOID(unsafe.Pointer(wnd)))
	if hWnd == 0 {
		return nil, errors.New(fmt.Sprintf("dion: window handle is empty ~> %x", kernel32.GetLastError()))
	}

	wnd.hWnd = hWnd

	getD2D1Factory().CreateHwndRenderTarget(d2d1.RenderTargetProperties(), d2d1.HwndRenderTargetProperties(hWnd, width, height), &wnd.pRT)

	dionWindows[wnd] = struct{}{}

	user32.UpdateWindow(hWnd)
	user32.ShowWindow(hWnd, user32.SW_SHOW)

	return wnd, nil
}