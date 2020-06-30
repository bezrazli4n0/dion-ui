package user32

import (
	"syscall"
	"unsafe"
)

var (
	user32DLL = syscall.NewLazyDLL("user32.dll")

	pRegisterClassEx = user32DLL.NewProc("RegisterClassExW")
	pShowWindow = user32DLL.NewProc("ShowWindow")
	pUpdateWindow = user32DLL.NewProc("UpdateWindow")
	pGetMessage = user32DLL.NewProc("GetMessageW")
	pTranslateMessage = user32DLL.NewProc("TranslateMessage")
	pDispatchMessage = user32DLL.NewProc("DispatchMessageW")
	pCreateWindowEx = user32DLL.NewProc("CreateWindowExW")
	pDefWindowProc = user32DLL.NewProc("DefWindowProcW")
	pLoadCursor = user32DLL.NewProc("LoadCursorW")
	pPostQuitMessage = user32DLL.NewProc("PostQuitMessage")
)

// RegisterClassEx регистрирует окно в системе
func RegisterClassEx(wc *WNDCLASSEXW) bool {
	ret, _, _ := pRegisterClassEx.Call(uintptr(unsafe.Pointer(wc)))

	if ret == 0 {
		return false
	}

	return true
}

// ShowWindow устанавливает специальный 'state' окна
func ShowWindow(hWnd HWND, nCmdShow int32) bool {
	ret, _, _ := pShowWindow.Call(uintptr(hWnd), uintptr(nCmdShow))

	if ret == 0 {
		return false
	}

	return true
}

// UpdateWindow обновляет клиентскую часть окна
func UpdateWindow(hWnd HWND) bool {
	ret, _, _ := pUpdateWindow.Call(uintptr(hWnd))

	if ret == 0 {
		return false
	}

	return true
}

// GetMessage получает сообщение из очереди сообщений вызывающего потока
func GetMessage(lpMsg *MSG, hWnd HWND, wMsgFilterMin, wMsgFilterMax uint32) bool {
	ret, _, _ := pGetMessage.Call(uintptr(unsafe.Pointer(lpMsg)), uintptr(hWnd), uintptr(wMsgFilterMin), uintptr(wMsgFilterMax))

	if ret == 0 {
		return false
	}

	return true
}

// TranslateMessage транслирует сообщения
func TranslateMessage(lpMsg *MSG) bool {
	ret, _, _ := pTranslateMessage.Call(uintptr(unsafe.Pointer(lpMsg)))

	if ret == 0 {
		return false
	}

	return true
}

// DispatchMessage отправляет сообщения в оконную процедуру
func DispatchMessage(lpMsg *MSG) LRESULT {
	ret, _, _ := pDispatchMessage.Call(uintptr(unsafe.Pointer(lpMsg)))
	return LRESULT(ret)
}

// DefWindowProc обрабатывает оконное сообщение по умолчанию
func DefWindowProc(hWnd HWND, uMsg uint32, wParam WPARAM, lParam LPARAM) LRESULT {
	ret, _, _ := pDefWindowProc.Call(uintptr(hWnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam))
	return LRESULT(ret)
}

// LoadCursor загружает курсор
func LoadCursor(lpCursorName uint32) HCURSOR {
	ret, _, _ := pLoadCursor.Call(uintptr(0), uintptr(lpCursorName))
	return HCURSOR(ret)
}

// PostQuitMessage указывает системе, что поток должен быть завершен
func PostQuitMessage(nExitCode int) {
	pPostQuitMessage.Call(uintptr(nExitCode))
}

// CreateWindowEx создает окно
func CreateWindowEx(dwExStyle DWORD, lpClassName, lpWindowName *uint16, dwStyle DWORD, X, Y, nWidth, nHeight int32, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE) HWND {
	ret, _, _ := pCreateWindowEx.Call(
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(lpClassName)),
		uintptr(unsafe.Pointer(lpWindowName)),
		uintptr(dwStyle),
		uintptr(X),
		uintptr(Y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(0),
	)
	return HWND(ret)
}