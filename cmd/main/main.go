package main

import (
	"fmt"
	"gateway/models"
	"gateway/pkg/control"
	"gateway/pkg/view"

	"fyne.io/fyne/v2"
)

type main_window struct {
	list_panel *fyne.Container
	text_panel *fyne.Container
	body_panel *fyne.Container
	main_lay   *fyne.Container
}

var list_panel *fyne.Container

func main() {
	models.FileName = ""
	models.Init()
	if control.CheckARGS() {
		fmt.Println("OPEN FILE ARGS")
	}
	view.MainWindowInit()
}
