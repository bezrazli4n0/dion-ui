package dion

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// windowMarkup содержит разметку окна
type windowMarkup struct {
	XMLName xml.Name `xml:"Window"`
	Title string `xml:"title,attr"`
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
	Width int `xml:"width,attr"`
	Height int `xml:"height,attr"`
	Autoclose bool `xml:"autoclose,attr"`

	OnLMouseDownCallback string `xml:"onLMouseDown,attr"`
	OnLMouseUpCallback string `xml:"onLMouseUp,attr"`
	OnMouseMove string `xml:"onMouseMove,attr"`
}

// WindowState внешний интерфейс взаимодействия с движком
type WindowState interface {
	GetWindow() *Window

	LoadUIFromFile(filePath string)
	LoadUIFromFileWithInterval(filePath string, interval time.Duration)
}

// windowStateImpl содержит состояние окна
type windowStateImpl struct {
	loadedWindow Window
	loadedWindowMarkup string
	loadedWindowFunctions map[string]interface{}
}

// GetWindow возвращает окно
func (w *windowStateImpl) GetWindow() *Window {
	return &w.loadedWindow
}

// NewWindowEngine иниициализирует движок для декларативного создания окна
func NewWindowEngine(windowFunctions map[string]interface{}) (WindowState, error) {
	wnd, err := NewWindow("DionUI", 0, 0, 640, 480)

	windowState := &windowStateImpl{
		loadedWindow: wnd,
		loadedWindowFunctions: windowFunctions,
	}

	return windowState, err
}

// LoadUIFromFile загружает интерфейс из файла
func (w *windowStateImpl) LoadUIFromFile(filePath string) {
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	if w.loadedWindowMarkup == string(data) {
		log.Println("NO UPDATE")
		return
	}

	windowMarkupState := windowMarkup{}
	xml.Unmarshal(data, &windowMarkupState)

	log.Println("UPDATE WINDOW")

	if w.loadedWindow.GetTitle() != windowMarkupState.Title {
		w.loadedWindow.SetTitle(windowMarkupState.Title)
	}

	width, height := w.loadedWindow.GetSize()

	if width != windowMarkupState.Width {
		w.loadedWindow.SetSize(windowMarkupState.Width, height)
		width = windowMarkupState.Width
	}

	if height != windowMarkupState.Height {
		w.loadedWindow.SetSize(width, windowMarkupState.Height)
		height = windowMarkupState.Height
	}

	x, y := w.loadedWindow.GetPos()

	if x != windowMarkupState.X {
		w.loadedWindow.SetPos(windowMarkupState.X, y)
		x = windowMarkupState.X
	}

	if y != windowMarkupState.Y {
		w.loadedWindow.SetPos(x, windowMarkupState.Y)
		y = windowMarkupState.Y
	}

	if windowMarkupState.Autoclose {
		w.loadedWindow.AttachCallback(OnClose, func() {
			w.loadedWindow.Close()
		})
	}

	// Callback's
	if callbackFunc, ok := w.loadedWindowFunctions[windowMarkupState.OnLMouseDownCallback]; ok {
		w.loadedWindow.AttachCallback(OnLMouseButtonDown, callbackFunc)
	}

	if callbackFunc, ok := w.loadedWindowFunctions[windowMarkupState.OnLMouseUpCallback]; ok {
		w.loadedWindow.AttachCallback(OnLMouseButtonUp, callbackFunc)
	}

	if callbackFunc, ok := w.loadedWindowFunctions[windowMarkupState.OnMouseMove]; ok {
		w.loadedWindow.AttachCallback(OnMouseMove, callbackFunc)
	}

	w.loadedWindowMarkup = string(data)
}

// LoadUIFromFileWithInterval загружает интерфейс из файла каждый интервал времени
func (w *windowStateImpl) LoadUIFromFileWithInterval(filePath string, interval time.Duration) {
	go func() {
		for {
			w.LoadUIFromFile(filePath)
			time.Sleep(interval)
		}
	}()
}