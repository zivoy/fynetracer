package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
	"renderShower/helpers"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var (
	width  = 800
	height = 800
)
var buffer *image.RGBA
var raster *canvas.Raster

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	buffer = image.NewRGBA(image.Rectangle{Max: image.Point{X: width, Y: height}})

	// go func() {
	// 	sleepAmount := time.Second / 10
	// 	for i := 0; i < 3000; i++ {
	// 		setRandomPixel()
	// 		if i%10 == 0 {
	// 			raster.Refresh()
	// 		}
	// 		time.Sleep(sleepAmount)
	// 	}
	// }()

	// rect := canvas.NewRectangle(color.White)
	// rect.SetMinSize(fyne.NewSize(150, 100))
	raster = canvas.NewRasterWithPixels(
		func(x, y, w, h int) color.Color {
			xC := x * width / w
			yC := y * height / h
			// use nfnt/resize
			return buffer.At(xC, yC)
		})
	// raster = canvas.NewRasterFromImage(buffer)
	raster.SetMinSize(fyne.NewSize(300, 300))

	timeTook := widget.NewLabel("last render:")
	timeTook.MinSize()

	var startTime, endTime time.Time
	timeRender := func() {
		startTime = time.Now()
		Render()
		endTime = time.Now()
		timeTook.SetText(fmt.Sprintf("last render: %s", endTime.Sub(startTime).String()))
	}
	action := widget.NewButton("Render", timeRender)
	timer := time.NewTicker(time.Second / 30)
	var stop chan interface{}
	autoRender := widget.NewCheck("Auto Render", func(b bool) {
		if b {
			stop = make(chan interface{})
			go func() {
				s := false
				for {
					if s {
						break
					}
					select {
					case <-timer.C:
						timeRender()
					case <-stop:
						s = true
					}
				}
			}()
		} else {
			// timer.Stop()
			close(stop)
		}
	})
	autoRender.SetChecked(true)

	spin := widget.NewSlider(0, 360)
	pitch := widget.NewSlider(-90, 90)
	const deg2rad = math.Pi / 180
	updateLight := func(_ float64) {
		aX := spin.Value * deg2rad
		aY := pitch.Value * deg2rad

		lightDir.X = -math.Cos(aX)*math.Cos(aY) + spherePos.X
		lightDir.Y = -math.Sin(aX)*math.Cos(aY) + spherePos.Y
		lightDir.Z = -math.Sin(aY) + spherePos.Z
		lightDir.Normelize()
	}
	spin.OnChanged = updateLight
	pitch.OnChanged = updateLight
	spin.SetValue(222)
	pitch.SetValue(-55)
	spin.OnChanged(spin.Value)
	pitch.OnChanged(pitch.Value)
	lightpos := container.New(layout.NewVBoxLayout(), spin, pitch) // widget.NewButton("print vlaues", func() {
	// 	fmt.Printf("spin.Value: %v\n", spin.Value)
	// 	fmt.Printf("pitch.Value: %v\n", pitch.Value)
	// 	fmt.Printf("lightDir: %v\n", lightDir)
	// })

	colPic := dialog.NewColorPicker("sphere colour", "select colour", func(c color.Color) {
		sphereCol = helpers.ColourFromColor(c)
	}, w)
	colourPicker := widget.NewButton("sphere colour", func() {
		colPic.Show()
	})
	colPic.Advanced = true
	colPic.Refresh()
	colPic.SetColor(sphereCol.RGBA())
	panel := container.New(layout.NewVBoxLayout(), timeTook, widget.NewLabel("------------------------------------------"), action, autoRender, lightpos, colourPicker)

	container := container.NewBorder(nil, nil, nil, panel, raster)
	w.SetContent(container)
	w.ShowAndRun()
}
