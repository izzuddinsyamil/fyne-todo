package main

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// init sqlite db
	// db, err := sql.Open("sqlite3", "./todo.db")
	// if err != nil {
	// 	log.Println("error when opening sqlite3 database", err)
	// }
	// _ = db

	a := app.New()
	w := a.NewWindow("Todo App")

	todoEntry := widget.NewEntry()
	todoEntry.SetPlaceHolder("Add your new todo")

	var todos []string
	submitTodoButton := widget.NewButton("Submit", func() {
		entry := todoEntry.Text
		if entry != "" {
			todos = append(todos, todoEntry.Text)
		}
	})

	list := widget.NewList(
		func() int {
			return len(todos)
		},
		func() fyne.CanvasObject {
			return widget.NewButton("", func() {})
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			text := fmt.Sprintf("%d. %s", i+1, todos[i])
			o.(*widget.Button).SetText(text)
		})

	if a.Driver().Device().IsMobile() {
		w.SetContent(renderMobileLayout())
	} else {
		w.SetContent(renderDesktopLayout(list, todoEntry, submitTodoButton))
	}

	w.Resize(fyne.NewSize(500, 200))
	w.ShowAndRun()
}

func renderMobileLayout() *fyne.Container {
	return nil
}

func renderDesktopLayout(list *widget.List, newTodo *widget.Entry, submit *widget.Button) *container.Split {
	return container.NewHSplit(list, container.NewVBox(
		newTodo,
		submit,
		layout.NewSpacer(),
	))
}
