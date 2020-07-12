package user32

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"syscall"
)

// HWND содержит дескриптор окна
type HWND syscall.Handle

// HINSTANCE указатель на объект ядра
type HINSTANCE syscall.Handle

// HICON дескриптор иконки
type HICON syscall.Handle

// HCURSOR дескриптор курсора
type HCURSOR syscall.Handle

// HBRUSH указывает на объект кисти
type HBRUSH syscall.Handle

// HMENU дескриптор меню объекта
type HMENU syscall.Handle

// HDC дескриптор девайса
type HDC syscall.Handle

const (
	CS_HREDRAW = 0x0002
	CS_VREDRAW = 0x0001

	SW_MAXIMIZE = 3
	SW_MINIMIZE = 6
	SW_RESTORE = 9
	SW_SHOW = 5
	SW_SHOWDEFAULT = 10
	SW_HIDE = 0

	TME_LEAVE = 0x00000002

	QS_ALLEVENTS = 0x04BF

	WS_OVERLAPPED = 0x00000000
	WS_SYSMENU = 0x00080000
	WS_CAPTION = 0x00C00000
	WS_THICKFRAME = 0x00040000
	WS_MINIMIZEBOX = 0x00020000
	WS_MAXIMIZEBOX = 0x00010000
	WS_OVERLAPPEDWINDOW = 0x00CF0000

	IDC_ARROW = 32512
	IDI_APPLICATION = 32512

	WM_CLOSE = 0x0010
	WM_LBUTTONDOWN = 0x0201
	WM_LBUTTONUP = 0x0202
	WM_NCCREATE = 0x0081
	WM_RBUTTONDOWN = 0x0204
	WM_RBUTTONUP = 0x0205
	WM_DESTROY = 0x0002
	WM_MOUSEMOVE = 0x0200
	WM_DISPLAYCHANGE = 0x007E
	WM_PAINT = 0x000F
	WM_SIZE = 0x0005
	WM_SIZING = 0x0214
	WM_MOUSELEAVE = 0x02A3
	WM_LBUTTONDBLCLK = 0x0203

	COLOR_WINDOW = 5

	GWLP_USERDATA = -21

	HWND_BOTTOM = 1
	HWND_NOTOPMOST = -2
	HWND_TOP = 0
	HWND_TOPMOST = -1

	SWP_ASYNCWINDOWPOS = 0x4000
	SWP_DEFERERASE = 0x2000
	SWP_DRAWFRAME = 0x0020
	SWP_FRAMECHANGED = 0x0020
	SWP_HIDEWINDOW = 0x0080
	SWP_NOACTIVATE = 0x0010
	SWP_NOCOPYBITS = 0x0100
	SWP_NOMOVE = 0x0002
	SWP_NOOWNERZORDER = 0x0200
	SWP_NOREDRAW = 0x0008
	SWP_NOREPOSITION = 0x0200
	SWP_NOSENDCHANGING = 0x0400
	SWP_NOSIZE = 0x0001
	SWP_NOZORDER = 0x0004
	SWP_SHOWWINDOW = 0x0040

	PM_REMOVE = 0x0001
)

type WNDCLASSEXW struct {
	CbSize        uint32
	Style         uint32
	LpfnWndProc   uintptr
	CbClsExtra    int32
	CbWndExtra    int32
	HInstance     HINSTANCE
	HIcon         HICON
	HCursor       HCURSOR
	HbrBackground HBRUSH
	LpszMenuName  *uint16
	LpszClassName *uint16
	HIconSm       HICON
}

type POINT struct {
	X int32
	Y int32
}

type MSG struct {
	Hwnd     HWND
	Message  uint32
	WParam   winapi.WPARAM
	LParam   winapi.LPARAM
	Time     winapi.DWORD
	Pt       POINT
	lPrivate winapi.DWORD
}

type CREATESTRUCTW struct {
	LpCreateParams winapi.LPVOID
	HInstance HINSTANCE
	HMenu HMENU
	HwndParent HWND
	Cy, Cx, Y, X int32
	Style int32
	LpszName *uint16
	LpszClass *uint16
	DwExStyle winapi.DWORD
}

type RECT struct {
	Left, Top, Right, Bottom int32
}

type PAINTSTRUCT struct {
	Hdc HDC
	FErase winapi.BOOL
	RcPaint RECT
	FRestore winapi.BOOL
	FIncUpdate winapi.BOOL
	RgbReserved [32]winapi.BYTE
}

type TRACKMOUSEEVENT struct {
	CbSize winapi.DWORD
	DwFlags winapi.DWORD
	HwndTrack HWND
	DwHoverTime winapi.DWORD
}