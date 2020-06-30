package user32

import (
	"syscall"
)

// HWND содержит дескриптор окна
type HWND syscall.Handle

// Handle указатель на объект ядра
type Handle syscall.Handle

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

// DWORD 4 байта целое
type DWORD uint32

// LPARAM указатель 4 байта
type LPARAM uintptr

// WPARAM указатель 4 байта
type WPARAM uintptr

// LRESULT указатель 4 байта
type LRESULT uintptr

const (
	CS_HREDRAW = 0x0002
	CS_VREDRAW = 0x0001

	SW_MAXIMIZE = 3
	SW_MINIMIZE = 6
	SW_RESTORE = 9
	SW_SHOW = 5
	SW_SHOWDEFAULT = 10
	SW_HIDE = 0

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

	COLOR_WINDOW = 5
)

type WNDCLASSEXW struct {
	CbSize uint32
	Style uint32
	LpfnWndProc uintptr
	CbClsExtra int32
	CbWndExtra int32
	HInstance HINSTANCE
	HIcon HICON
	HCursor HCURSOR
	HbrBackground HBRUSH
	LpszMenuName *uint16
	LpszClassName *uint16
	HIconSm HICON
}

type POINT struct {
	X int32
	Y int32
}

type MSG struct {
	Hwnd HWND
	Message uint32
	WParam WPARAM
	LParam LPARAM
	Time DWORD
	Pt POINT
	lPrivate DWORD
}