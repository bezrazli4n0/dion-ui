package dion

import "github.com/bezrazli4n0/dion-ui/internal/pkg/winapi/d2d1"

var pD2D1Factory *d2d1.ID2D1Factory

func initGraphics() {
	d2d1.CreateFactory(d2d1.FACTORY_TYPE_SINGLE_THREADED, &pD2D1Factory)
}

func getD2D1Factory() *d2d1.ID2D1Factory {
	return pD2D1Factory
}

func freeGraphics() {
	pD2D1Factory.Release()
}