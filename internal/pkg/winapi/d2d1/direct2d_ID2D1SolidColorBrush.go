package d2d1

// ID2D1SolidColorBrush отрисовывает область сплошным цветом
type ID2D1SolidColorBrush struct {
	ID2D1Brush
}

// vtblID2D1SolidColorBrush виртуальная таблица для ID2D1SolidColorBrush
type vtblID2D1SolidColorBrush struct {
	vtblID2D1Brush
	SetColor uintptr
	GetColor uintptr
}