package dion

import "testing"

func TestNewWindow (t *testing.T) {
	Init()

	wnd, err := NewWindow("Window", 0, 0, 640, 480)
	if err != nil {
		t.Error(err)
	}

	wnd.SetTitle("Window-test")
	if wnd.GetTitle() != "Window-test" {
		t.Error("Window title must be changed")
	}

	Run()
}
