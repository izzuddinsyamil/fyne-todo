package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Todo App")

	title := canvas.NewText("Todo App", color.White)
	titleContainer := container.New(layout.NewCenterLayout(), title)

	todoEntry := widget.NewEntry()
	todoEntry.SetPlaceHolder("Add your new todo")

	submitTodoButton := widget.NewButton("Submit", func() {})

	todoEntryContainer := container.New(
		layout.NewGridLayout(2),
		todoEntry,
		submitTodoButton,
	)

	w.SetContent(container.NewAdaptiveGrid(1,
		titleContainer,
		todoEntryContainer,
	))

	w.Resize(fyne.NewSize(100, 100))
	w.ShowAndRun()
}
