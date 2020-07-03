package d2d1

// ID2D1Brush содержит базовый интерфейс кисти
type ID2D1Brush struct {
	ID2D1Resource
}

// vtblID2D1Brush виртуальная таблица для ID2D1Brush
type vtblID2D1Brush struct {
	vtblID2D1Resource
	SetOpacity uintptr
	SetTransform uintptr
	GetOpacity uintptr
	GetTransform uintptr
}