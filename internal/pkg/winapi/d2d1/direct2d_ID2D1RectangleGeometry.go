package d2d1

// ID2D1RectangleGeometry
type ID2D1RectangleGeometry struct {
	ID2D1Geometry
}

// vtblID2D1RectangleGeometry виртуальная таблица для ID2D1RectangleGeometry
type vtblID2D1RectangleGeometry struct {
	vtblID2D1Geometry
	GetRect uintptr
}
