package com

import (
	"syscall"
	"unsafe"
)

// IUnknown базовый COM интерфейс
type IUnknown struct {
	Vtbl unsafe.Pointer
}

// vtblIUnknown виртуальная таблица для IUnknown
type VtblIUnknown struct {
	QueryInterface uintptr
	AddRef uintptr
	Release uintptr
}

// vmt возвращает указатель на виртуальную таблицу
func (obj *IUnknown) vmt() *VtblIUnknown {
	return (*VtblIUnknown)(obj.Vtbl)
}

// QueryInterface https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-queryinterface(refiid_void)
func (obj *IUnknown) QueryInterface(guid *Guid) unsafe.Pointer {
	var dest unsafe.Pointer
	syscall.Syscall(obj.vmt().QueryInterface, 3, uintptr(unsafe.Pointer(obj)), uintptr(unsafe.Pointer(guid)), uintptr(unsafe.Pointer(&dest)))
	return dest
}

// AddRef https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-addref
func (obj *IUnknown) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(obj.vmt().AddRef, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return uint32(ret)
}

// Release https://docs.microsoft.com/en-us/windows/win32/api/unknwn/nf-unknwn-iunknown-release
func (obj *IUnknown) Release() uint32 {
	ret, _, _ := syscall.Syscall(obj.vmt().Release, 1, uintptr(unsafe.Pointer(obj)), 0, 0)
	return uint32(ret)
}