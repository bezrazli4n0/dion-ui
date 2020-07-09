package kernel32

import "syscall"

const (
	WAIT_TIMEOUT = 0x00000102
)

// HMODULE дескриптор модуля
type HMODULE syscall.Handle