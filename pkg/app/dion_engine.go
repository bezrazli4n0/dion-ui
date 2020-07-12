package dion

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// windowCanvas содержит canvas в окне
type windowCanvas struct {
	CanvasObjects []windowCanvasObject `xml:",any"`
}

// windowCanvasObject содержит canvas примитив
type windowCanvasObject struct {
	XMLName xml.Name
	X float32 `xml:"x,attr"`
	Y float32 `xml:"y,attr"`
	Width string `xml:"width,attr"`
	Height string `xml:"height,attr"`
	Color string `xml:"color,attr"`
	StrokeWidth float32 `xml:"stroke,attr"`
}

// windowMarkup содержит разметку окна
type windowMarkup struct {
	XMLName xml.Name `xml:"Window"`
	Title string `xml:"title,attr"`
	X int `xml:"x,attr"`
	Y int `xml:"y,attr"`
	Width int `xml:"width,attr"`
	Height int `xml:"height,attr"`
	Autoclose bool `xml:"autoclose,attr"`
	BackgroundColor string `xml:"backgroundColor,attr"`

	OnLMouseDownCallback string `xml:"onLMouseDown,attr"`
	OnLMouseUpCallback string `xml:"onLMouseUp,attr"`
	OnMouseMove string `xml:"onMouseMove,attr"`

	OnResize string `xml:"onResize,attr"`

	OnRMouseDownCallback string `xml:"onRMouseDown,attr"`
	OnRMouseUpCallback string `xml:"onRMouseUp,attr"`

	Canvas windowCanvas `xml:"Canvas"`
	LayoutWidget layoutWidget `xml:",any"`
}

// layoutWidget содержит корневой виджет окна
type layoutWidget struct {
	XMLName xml.Name

	X float32 `xml:"x,attr"`
	Y float32 `xml:"y,attr"`
	Width float32 `xml:"width,attr"`
	Height float32 `xml:"height,attr"`
	FontName string `xml:"fontName,attr"`
	FontSize float32 `xml:"fontSize,attr"`
	WidgetColor string `xml:"color,attr"`
	InnerArg string `xml:",chardata"`
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
	wnd, err := NewWindow("DionUI", 0, 0, 640, 480, nil)

	windowState := &windowStateImpl{
		loadedWindow: wnd,
		loadedWindowFunctions: windowFunctions,
	}

	return windowState, err
}

// parseColor парсит цвет из строки
func parseColor(color string) (r, g, b, a byte) {
	color = strings.ReplaceAll(color, " ", "")

	if strings.Contains(color, "#") {
		// #FF0000
		rVal, _ := strconv.ParseUint(color[1:3], 16, 8)
		gVal, _ := strconv.ParseUint(color[3:5], 16, 8)
		bVal, _ := strconv.ParseUint(color[5:], 16, 8)
		r = byte(rVal)
		g = byte(gVal)
		b = byte(bVal)
		a = 255
	} else if strings.Contains(color, ",") {
		parsedSlice := strings.Split(color, ",")
		rVal, _ := strconv.Atoi(parsedSlice[0])
		gVal, _ := strconv.Atoi(parsedSlice[1])
		bVal, _ := strconv.Atoi(parsedSlice[2])
		if len(parsedSlice) == 4 {
			aVal, _ := strconv.Atoi(parsedSlice[3])
			a = byte(aVal)
		} else {
			a = 255
		}
		r = byte(rVal)
		g = byte(gVal)
		b = byte(bVal)
	} else {
		r = 255
		g = 255
		b = 255
		a = 255
	}
	return
}

// setCanvasObjectSize парсит и устанавливает размер canvas примитива
func setCanvasObjectSize(object CanvasObject, width, height string) {
	widthPercent := 0.0
	heightPercent :=  0.0

	widthValue := 0.0
	heightValue := 0.0

	var err error

	if strings.Contains(width, "%") {
		width = strings.ReplaceAll(width, "%", "")
		widthPercent, err = strconv.ParseFloat(width, 32)
		if err != nil {
			widthPercent = -1.0
		}
	} else {
		widthValue, err = strconv.ParseFloat(width, 32)
		if err != nil {
			widthValue = -1.0
		}
	}

	if strings.Contains(height, "%") {
		height = strings.ReplaceAll(height, "%", "")
		heightPercent, err = strconv.ParseFloat(height, 32)
		if err != nil {
			heightPercent = -1.0
		}
	} else {
		heightValue, err = strconv.ParseFloat(height, 32)
		if err != nil {
			heightValue = -1.0
		}
	}

	if widthPercent != -1.0 || heightPercent != -1.0 {
		object.setPercentSize(float32(widthPercent), float32(heightPercent))
	}

	object.SetSize(float32(widthValue), float32(heightValue))
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

	/* State manager */
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
			Exit()
		})
	}

	// Callback's
	/* Left mouse button */
	if callbackFunc, ok := w.loadedWindowFunctions[windowMarkupState.OnLMouseDownCallback]; ok {
		w.loadedWindow.AttachCallback(OnLMouseButtonDown, callbackFunc)
	}

	if callbackFunc, ok := w.loadedWindowFunctions[windowMarkupState.OnLMouseUpCallback]; ok {
		w.loadedWindow.AttachCallback(OnLMouseButtonUp, callbackFunc)
	}

	if callbackFunc, ok := w.loadedWindowFunctions[windowMarkupState.OnMouseMove]; ok {
		w.loadedWindow.AttachCallback(OnMouseMove, callbackFunc)
	}

	/* Right mouse button */
	if callbackFunc, ok := w.loadedWindowFunctions[windowMarkupState.OnRMouseDownCallback]; ok {
		w.loadedWindow.AttachCallback(OnRMouseButtonDown, callbackFunc)
	}

	if callbackFunc, ok := w.loadedWindowFunctions[windowMarkupState.OnRMouseUpCallback]; ok {
		w.loadedWindow.AttachCallback(OnRMouseButtonUp, callbackFunc)
	}

	// Background color
	r, g, b, a := parseColor(windowMarkupState.BackgroundColor)
	w.loadedWindow.SetBackgroundColor(r, g, b, a)

	// Resize
	if callbackFunc, ok := w.loadedWindowFunctions[windowMarkupState.OnResize]; ok {
		w.loadedWindow.AttachCallback(OnResize, callbackFunc)
	}

	// Canvas state
	canvas := []CanvasObject{}

	// TODO: add more canvas primitives
	for _, obj := range windowMarkupState.Canvas.CanvasObjects {
		r, g, b, a = parseColor(obj.Color)
		color := Color{r, g, b, a}

		switch obj.XMLName.Local {
		case "FillRectangle":
			fillRect := NewRectangle(obj.X, obj.Y, 0.0, 0.0, true, Color{}, obj.StrokeWidth)
			setCanvasObjectSize(fillRect, obj.Width, obj.Height)
			fillRect.SetColorRGBA(color)
			canvas = append(canvas, fillRect)
			break

		case "DrawRectangle":
			drawRect := NewRectangle(obj.X, obj.Y, 0.0, 0.0, false, Color{}, obj.StrokeWidth)
			setCanvasObjectSize(drawRect, obj.Width, obj.Height)
			drawRect.SetColorRGBA(color)
			canvas = append(canvas, drawRect)
			break
		}
	}

	w.loadedWindow.SetCanvas(&Canvas{canvas})

	// Load widget
	// TODO: imporve layout system
	switch windowMarkupState.LayoutWidget.XMLName.Local {
	case "Label":
		r, g, b, a := parseColor(windowMarkupState.LayoutWidget.WidgetColor)
		lblColor := Color{r, g, b, a}
		lbl := NewLabel(strings.TrimSpace(windowMarkupState.LayoutWidget.InnerArg), windowMarkupState.LayoutWidget.FontName, windowMarkupState.LayoutWidget.X, windowMarkupState.LayoutWidget.Y, windowMarkupState.LayoutWidget.Width, windowMarkupState.LayoutWidget.Height, windowMarkupState.LayoutWidget.FontSize, lblColor)
		w.loadedWindow.SetWidget(lbl)
	}

	w.loadedWindowMarkup = string(data)
}

// LoadUIFromFileWithInterval загружает интерфейс из файла каждый интервал времени
func (w *windowStateImpl) LoadUIFromFileWithInterval(filePath string, interval time.Duration) {
	go func() {
		for !w.loadedWindow.(*window).isClosed {
			w.LoadUIFromFile(filePath)
			time.Sleep(interval)
		}
	}()
}