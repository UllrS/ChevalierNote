package main

import (
	"fmt"
	"gateway/pkg/docmanager"
	"gateway/pkg/view"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type main_window struct {
	list_panel *fyne.Container
	text_panel *fyne.Container
	body_panel *fyne.Container
	main_lay   *fyne.Container
}

var list_panel *fyne.Container

func main() {
	fmt.Println("DDD")
	docmanager.ReadDoc("data.json")
	icon, _ := fyne.LoadResourceFromPath("assets/image/knight.png")
	a := app.New()
	w := a.NewWindow("Note")
	w.SetIcon(icon)
	w.Resize(fyne.NewSize(680, 420))

	sd := widget.NewLabel("sda")
	sd.Text = "sadfg"
	d := widget.NewSeparator()
	control_panel := view.ControlPanelCreate()
	body_panel := view.BodyCreate()
	main_lay := container.NewBorder(control_panel, nil, nil, nil, d, body_panel) //(widjets.ControlButton("1"), nil, nil, nil)

	w.SetContent(main_lay)

	w.ShowAndRun()
}
