package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello")
	_ = w

	bAC := widget.NewButton("AC", func() {
		log.Println("AC basıldı")
	})
	bNeg := widget.NewButton("+/-", func() {
		log.Println("Negatif sayı basıldı")
	})

	b4 := widget.NewButton("b4", func() {
		log.Println("b4 basıldı")

	})

	t1 := canvas.NewText("1", color.White)

	bg := container.New(layout.NewAdaptiveGridLayout(4), bAC, bNeg, btnYuzde(), btnBolme(), b4)

	// g := container.New(layout.NewAdaptiveGridLayout(1), t1, b1, b2, b3, b4)
	g := container.New(layout.NewGridLayoutWithRows(2), t1, bg)

	w.SetContent(g)
	w.ShowAndRun()

}

func turuncu() *canvas.Rectangle {
	btnColor := canvas.NewRectangle(
		color.NRGBA{R: 255, G: 159, B: 11, A: 255})
	return btnColor
}

func btnYuzde() *fyne.Container {
	btn := widget.NewButton("%", nil)

	container1 := container.New(
		layout.NewMaxLayout(),
		btn,
	)
	return container1
}

func btnBolme() *fyne.Container {
	btn := widget.NewButton("/", nil)

	container1 := container.New(
		layout.NewMaxLayout(),
		turuncu(),
		btn,
	)
	return container1
}
