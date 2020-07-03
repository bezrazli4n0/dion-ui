package dion

import (
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"
	"github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/dwrite"
)

var pD2D1Factory *d2d1.ID2D1Factory
var pDWriteFactory *dwrite.IDWriteFactory

var dpiX, dpiY float32

func initGraphics() {
	d2d1.CreateFactory(d2d1.FACTORY_TYPE_SINGLE_THREADED, &pD2D1Factory)
	dwrite.CreateFactory(dwrite.FACTORY_TYPE_ISOLATED, &pDWriteFactory)

	dpiX, dpiY = pD2D1Factory.GetDesktopDpi()
}

func getD2D1Factory() *d2d1.ID2D1Factory {
	return pD2D1Factory
}

func getDWriteFactory() *dwrite.IDWriteFactory {
	return pDWriteFactory
}

func freeGraphics() {
	pDWriteFactory.Release()
	pD2D1Factory.Release()
}