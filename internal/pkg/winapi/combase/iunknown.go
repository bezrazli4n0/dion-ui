package combase

// IUnknown базовый COM интерфейс
type IUnknown struct {
	QueryInterface uintptr
	AddRef uintptr
	Release uintptr
}