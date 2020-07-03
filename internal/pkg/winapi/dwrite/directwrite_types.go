package dwrite

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/com"

type FACTORY_TYPE int32

const (
	FACTORY_TYPE_SHARED FACTORY_TYPE = iota
	FACTORY_TYPE_ISOLATED
)

var IID_IDWriteFactory com.Guid = com.Guid{0xb859ee5a,0xd838,0x4b5b,[8]byte{0xa2, 0xe8, 0x1a, 0xdc, 0x7d, 0x93, 0xdb, 0x48}}