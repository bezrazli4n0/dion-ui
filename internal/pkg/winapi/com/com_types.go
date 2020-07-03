package com

// Guid статический уникальный идентификатор, используется COM интерфейсами
type Guid struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}
