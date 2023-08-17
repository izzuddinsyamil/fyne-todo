package main

import (
	"database/sql"
	"fmt"
	"image/color"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {

	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		log.Println("error when opening sqlite3 database", err)
	}

	fmt.Printf("db %v\n", db)

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
