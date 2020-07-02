package d2d1

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/combase"
)

type FACTORY_TYPE int32

const (
	FACTORY_TYPE_SINGLE_THREADED FACTORY_TYPE = 0
)

var IID_ID2D1Factory combase.Guid = combase.Guid{0x06152247,0x6f50,0x465a,[8]byte{0x92, 0x45, 0x11, 0x8b, 0xfd, 0x3b, 0x60, 0x07}}
