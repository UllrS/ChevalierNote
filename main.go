package main

import (
	"gateway/models"
	"gateway/pkg/control"
	"gateway/pkg/view"
	"os"

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
	if len(os.Args) > 1 {
		models.FileName = os.Args[1]
		control.DataLoad()
	}
	models.TreeItemMap = map[string]*models.TreeItem{}
	view.MainWindowInit()
}
