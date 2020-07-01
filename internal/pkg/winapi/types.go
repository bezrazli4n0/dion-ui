package winapi

import "syscall"

// DWORD 4 байта целое
type DWORD uint32

// LPARAM указатель 4 байта
type LPARAM uintptr

// WPARAM указатель 4 байта
type WPARAM uintptr

// LRESULT указатель 4 байта
type LRESULT uintptr

// HRESULT указатель 4 байта
type HRESULT uintptr

// Handle указатель на объект ядра
type Handle syscall.Handle

// LONG_PTR указатель
type LONG_PTR uintptr

// LPVOID указатель
type LPVOID uintptr

// PVOID указатель
type PVOID uintptr

// BOOL 4 байта
type BOOL int