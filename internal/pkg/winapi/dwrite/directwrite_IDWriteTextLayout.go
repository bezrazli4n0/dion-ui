package dwrite

type IDWriteTextLayout struct {
	IDWriteTextFormat
}

// vtblIDWriteTextLayout виртуальная таблица для IDWriteTextLayout
type vtblIDWriteTextLayout struct {
	vtblIDWriteTextFormat
	SetMaxWidth uintptr
	SetMaxHeight uintptr
	SetFontCollection uintptr
	SetFontFamilyName uintptr
	SetFontWeight uintptr
	SetFontStyle uintptr
	SetFontStretch uintptr
	SetFontSize uintptr
	SetUnderline uintptr
	SetStrikethrough uintptr
	SetDrawingEffect uintptr
	SetInlineObject uintptr
	SetTypography uintptr
	SetLocaleName uintptr
	GetMaxWidth uintptr
	GetMaxHeight uintptr
	GetFontCollection uintptr
	GetFontFamilyNameLength uintptr
	GetFontFamilyName uintptr
	GetFontWeight uintptr
	GetFontStyle uintptr
	GetFontStretch uintptr
	GetFontSize uintptr
	GetUnderline uintptr
	GetStrikethrough uintptr
	GetDrawingEffect uintptr
	GetInlineObject uintptr
	GetTypography uintptr
	GetLocaleNameLength uintptr
	GetLocaleName uintptr
	Draw uintptr
	GetLineMetrics uintptr
	GetMetrics uintptr
	GetOverhangMetrics uintptr
	GetClusterMetrics uintptr
	DetermineMinWidth uintptr
	HitTestPoint uintptr
	HitTestTextPosition uintptr
	HitTestTextRange uintptr
}