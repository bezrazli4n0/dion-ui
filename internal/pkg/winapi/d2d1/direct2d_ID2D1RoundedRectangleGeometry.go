package d2d1

// ID2D1RoundedRectangleGeometry
type ID2D1RoundedRectangleGeometry struct {
	ID2D1Geometry
}

// vtblID2D1RoundedRectangleGeometry виртуальная таблица для ID2D1RoundedRectangleGeometry
type vtblID2D1RoundedRectangleGeometry struct {
	vtblID2D1Geometry
	GetRoundedRect uintptr
}
