package d2d1

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi"
	"math"
	"syscall"
	"unsafe"
)

// ID2D1Geometry
type ID2D1Geometry struct {
	ID2D1Resource
}

// vtblID2D1Geometry виртуальная таблица для ID2D1Geometry
type vtblID2D1Geometry struct {
	vtblID2D1Resource
	GetBounds uintptr
	GetWidenedBounds uintptr
	StrokeContainsPoint uintptr
	FillContainsPoint uintptr
	CompareWithGeometry uintptr
	Simplify uintptr
	Tessellate uintptr
	CombineWithGeometry uintptr
	Outline uintptr
	ComputeArea uintptr
	ComputeLength uintptr
	ComputePointAtLength uintptr
	Widen uintptr
}

// vmt возвращает указатель на виртуальную таблицу
func (obj *ID2D1Geometry) vmt() *vtblID2D1Geometry {
	return (*vtblID2D1Geometry)(obj.Vtbl)
}

// FillContainsPoint https://docs.microsoft.com/en-us/windows/win32/api/d2d1/nf-d2d1-id2d1geometry-fillcontainspoint(d2d1_point_2f_constd2d1_matrix_3x2_f_float_bool)
func (obj *ID2D1Geometry) FillContainsPoint(point POINT_2F, worldTransform *MATRIX_3X2_F, flatteningTolerance float32) (winapi.HRESULT, bool) {
	var _contains winapi.BOOL = 0
	contains := false
	ret, _, _ := syscall.Syscall6(obj.vmt().FillContainsPoint, 5, uintptr(unsafe.Pointer(obj)), uintptr(*(*uint64)(unsafe.Pointer(&point))), uintptr(unsafe.Pointer(worldTransform)), uintptr(math.Float32bits(flatteningTolerance)), uintptr(unsafe.Pointer(&_contains)), 0)

	if _contains != 0 {
		contains = true
	}

	return winapi.HRESULT(ret), contains
}