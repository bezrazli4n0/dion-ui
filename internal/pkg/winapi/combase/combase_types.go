package combase

// Guid статический уникальный идентификатор, используется COM интерфейсами
type Guid struct {
	V1 int32
	V2 int16
	V3 int16
	V4 [8]byte
}
