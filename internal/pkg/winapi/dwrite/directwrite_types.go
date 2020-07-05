package dwrite

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/com"

type FACTORY_TYPE int32
type PARAGRAPH_ALIGNMENT int32
type TEXT_ALIGNMENT int32

// https://docs.microsoft.com/en-us/windows/win32/api/dwrite/ne-dwrite-dwrite_paragraph_alignment
const (
	PARAGRAPH_ALIGNMENT_NEAR PARAGRAPH_ALIGNMENT = iota
	PARAGRAPH_ALIGNMENT_FAR
	PARAGRAPH_ALIGNMENT_CENTER
)

// https://docs.microsoft.com/en-us/windows/win32/api/dwrite/ne-dwrite-dwrite_text_alignment
const (
	TEXT_ALIGNMENT_LEADING TEXT_ALIGNMENT = iota
	TEXT_ALIGNMENT_TRAILING
	TEXT_ALIGNMENT_CENTER
	TEXT_ALIGNMENT_JUSTIFIED
)

const (
	FACTORY_TYPE_SHARED FACTORY_TYPE = iota
	FACTORY_TYPE_ISOLATED
)

var IID_IDWriteFactory com.Guid = com.Guid{0xb859ee5a,0xd838,0x4b5b,[8]byte{0xa2, 0xe8, 0x1a, 0xdc, 0x7d, 0x93, 0xdb, 0x48}}