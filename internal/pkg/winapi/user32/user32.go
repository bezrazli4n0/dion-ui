package user32

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
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
	pSetWindowText = user32DLL.NewProc("SetWindowTextW")
	pGetWindowLongPtr = user32DLL.NewProc("GetWindowLongPtrW")
	pSetWindowLongPtr = user32DLL.NewProc("SetWindowLongPtrW")
	pUnregisterClass = user32DLL.NewProc("UnregisterClassW")
	pGetWindowRect = user32DLL.NewProc("GetWindowRect")
	pSetWindowPos = user32DLL.NewProc("SetWindowPos")
	pAdjustWindowRect = user32DLL.NewProc("AdjustWindowRect")
	pBeginPaint = user32DLL.NewProc("BeginPaint")
	pEndPaint = user32DLL.NewProc("EndPaint")
	pInvalidateRect = user32DLL.NewProc("InvalidateRect")
	pPeekMessage = user32DLL.NewProc("PeekMessageW")
	pDestroyWindow = user32DLL.NewProc("DestroyWindow")
	pMsgWaitForMultipleObjects = user32DLL.NewProc("MsgWaitForMultipleObjects")
	pSetCapture = user32DLL.NewProc("SetCapture")
	pReleaseCapture = user32DLL.NewProc("ReleaseCapture")
	pTrackMouseEvent = user32DLL.NewProc("TrackMouseEvent")
	pGetCursorPos = user32DLL.NewProc("GetCursorPos")
	pWindowFromPoint = user32DLL.NewProc("WindowFromPoint")
)

// WindowFromPoint получает дескриптор окна по координатам
func WindowFromPoint(Point POINT) HWND {
	ret, _, _ := pWindowFromPoint.Call(uintptr(Point.X), uintptr(Point.Y))
	return HWND(ret)
}

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
func DispatchMessage(lpMsg *MSG) winapi.LRESULT {
	ret, _, _ := pDispatchMessage.Call(uintptr(unsafe.Pointer(lpMsg)))
	return winapi.LRESULT(ret)
}

// DefWindowProc обрабатывает оконное сообщение по умолчанию
func DefWindowProc(hWnd HWND, uMsg uint32, wParam winapi.WPARAM, lParam winapi.LPARAM) winapi.LRESULT {
	ret, _, _ := pDefWindowProc.Call(uintptr(hWnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam))
	return winapi.LRESULT(ret)
}

// LoadCursor загружает курсор
func LoadCursor(lpCursorName uint32) HCURSOR {
	ret, _, _ := pLoadCursor.Call(uintptr(0), uintptr(lpCursorName))
	return HCURSOR(ret)
}

// TrackMouseEvent отслеживает различные события мыши в окне
func TrackMouseEvent(lpEventTrack *TRACKMOUSEEVENT) bool {
	ret, _, _ := pTrackMouseEvent.Call(uintptr(unsafe.Pointer(lpEventTrack)))

	if ret != 0 {
		return true
	}

	return false
}

// PostQuitMessage указывает системе, что поток должен быть завершен
func PostQuitMessage(nExitCode int) {
	pPostQuitMessage.Call(uintptr(nExitCode))
}

// SetWindowText устанавливает текст окна
func SetWindowText(hWnd HWND, lpString string) bool {
	_lpString, err := syscall.UTF16FromString(lpString)
	if err != nil {
		return false
	}

	ret, _, _ := pSetWindowText.Call(uintptr(hWnd), uintptr(unsafe.Pointer(&_lpString[0])))

	if ret == 0 {
		return false
	}

	return true
}

// GetWindowLongPtr получает информацию об окне
func GetWindowLongPtr(hWnd HWND, nIndex int) winapi.LONG_PTR {
	ret, _, _ := pGetWindowLongPtr.Call(uintptr(hWnd), uintptr(nIndex))
	return winapi.LONG_PTR(ret)
}

// SetWindowLongPtr устанавливает указанный атрибут окну
func SetWindowLongPtr(hWnd HWND, nIndex int, dwNewLong winapi.LONG_PTR) winapi.LONG_PTR {
	ret, _, _ := pSetWindowLongPtr.Call(uintptr(hWnd), uintptr(nIndex), uintptr(dwNewLong))
	return winapi.LONG_PTR(ret)
}

// Hiword возвращает старшую часть числа
func Hiword(val int32) int16 {
	return int16((val >> 16) & 0xffff)
}

// Loword возвращает младшую часть числа
func Loword(val int32) int16 {
	return int16(val & 0xffff)
}

// GetWindowRect возвращает размеры окна
func GetWindowRect(hWnd HWND, lpRect *RECT) bool {
	ret, _, _ := pGetWindowRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpRect)))

	if ret == 0 {
		return false
	}

	return true
}

// GetCursorPos возвращает позицию курсора на экране
func GetCursorPos(lpPoint *POINT) bool {
	ret, _, _ := pGetCursorPos.Call(uintptr(unsafe.Pointer(lpPoint)))

	if ret != 0 {
		return true
	}

	return false
}

// SetWindowPos изменяет размеры и позицию окна
func SetWindowPos(hWnd, hWndInsertAfter HWND, x, y, cx, cy int, uFlags uint32) bool {
	ret, _, _ := pSetWindowPos.Call(uintptr(hWnd), uintptr(hWndInsertAfter), uintptr(x), uintptr(y), uintptr(cx), uintptr(cy), uintptr(uFlags))

	if ret == 0 {
		return false
	}

	return true
}

// UnregisterClass высвобождает память окна
func UnregisterClass(lpClassName string, hInstance HINSTANCE) bool {
	_lpClassName, err := syscall.UTF16FromString(lpClassName)
	if err != nil {
		return false
	}

	ret, _, _ := pUnregisterClass.Call(uintptr(unsafe.Pointer(&_lpClassName[0])), uintptr(hInstance))

	if ret == 0 {
		return false
	}

	return true
}

// BeginPaint подготавливает окно для рисования
func BeginPaint(hWnd HWND, lpPaint *PAINTSTRUCT) HDC {
	ret, _, _ := pBeginPaint.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpPaint)))
	return HDC(ret)
}

// EndPaint сообщает окну о прекращении рисования
func EndPaint(hWnd HWND, lpPaint *PAINTSTRUCT) {
	pEndPaint.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpPaint)))
}

// SetCapture устанавливает захват мыши для определенного окна
func SetCapture(hWnd HWND) HWND {
	ret, _, _ := pSetCapture.Call(uintptr(hWnd))
	return HWND(ret)
}

// ReleaseCapture высвобождает захваченый курсор мыши
func ReleaseCapture() bool {
	ret, _, _ := pReleaseCapture.Call()

	res := winapi.BOOL(ret)
	if res != 0 {
		return true
	}

	return false
}

// AdjustWindowRect вычисляет необходимый размер окна
func AdjustWindowRect(lpRect *RECT, dwStyle winapi.DWORD, bMenu bool) bool {
	_bMenu := winapi.BOOL(0)
	if bMenu {
		_bMenu = winapi.BOOL(1)
	}

	ret, _, _ := pAdjustWindowRect.Call(uintptr(unsafe.Pointer(lpRect)), uintptr(dwStyle), uintptr(_bMenu))

	if ret == 0 {
		return false
	}

	return true
}

// InvalidateRect обновляет указанный регион
func InvalidateRect(hWnd HWND, lpRect *RECT, bErase bool) bool {
	_bErase := winapi.BOOL(0)
	if bErase {
		_bErase = winapi.BOOL(1)
	}
	ret, _, _ := pInvalidateRect.Call(uintptr(hWnd), uintptr(unsafe.Pointer(lpRect)), uintptr(_bErase))

	if ret == 0 {
		return false
	}

	return true
}

// DestroyWindow уничтожает окно
func DestroyWindow(hWnd HWND) bool {
	ret, _, _ := pDestroyWindow.Call(uintptr(hWnd))

	if ret == 0 {
		return  false
	}

	return true
}

// PeekMessage обрабатывает входящие сообщения окну
func PeekMessage(lpMsg *MSG, hWnd HWND, wMsgFilterMin, wMsgFilterMax, wRemoveMsg uint32) bool {
	ret, _, _ := pPeekMessage.Call(uintptr(unsafe.Pointer(lpMsg)), uintptr(hWnd), uintptr(wMsgFilterMin), uintptr(wMsgFilterMax), uintptr(wRemoveMsg))

	if ret == 0 {
		return false
	}

	return true
}

// MsgWaitForMultipleObjects ждёт объект ядра
func MsgWaitForMultipleObjects(nCount winapi.DWORD, pHandles *winapi.Handle, fWaitAll winapi.BOOL, dwMilliseconds, dwWakeMask winapi.DWORD) winapi.DWORD {
	ret, _, _ := pMsgWaitForMultipleObjects.Call(uintptr(nCount), uintptr(unsafe.Pointer(pHandles)), uintptr(fWaitAll), uintptr(dwMilliseconds), uintptr(dwWakeMask))
	return winapi.DWORD(ret)
}

// CreateWindowEx создает окно
func CreateWindowEx(dwExStyle winapi.DWORD, lpClassName, lpWindowName string, dwStyle winapi.DWORD, X, Y, nWidth, nHeight int32, hWndParent HWND, hMenu HMENU, hInstance HINSTANCE, lpParam winapi.LPVOID) HWND {
	_lpClassName, err := syscall.UTF16FromString(lpClassName)
	if err != nil {
		return 0
	}

	_lpWindowName, err := syscall.UTF16FromString(lpWindowName)
	if err != nil {
		return 0
	}

	ret, _, _ := pCreateWindowEx.Call(
		uintptr(dwExStyle),
		uintptr(unsafe.Pointer(&_lpClassName[0])),
		uintptr(unsafe.Pointer(&_lpWindowName[0])),
		uintptr(dwStyle),
		uintptr(X),
		uintptr(Y),
		uintptr(nWidth),
		uintptr(nHeight),
		uintptr(hWndParent),
		uintptr(hMenu),
		uintptr(hInstance),
		uintptr(lpParam),
	)
	return HWND(ret)
}