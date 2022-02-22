package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainWindowInit() {

	icon, _ := fyne.LoadResourceFromPath("assets/image/knight.png")
	a := app.New()
	w := a.NewWindow("Note")
	w.SetIcon(icon)
	w.Resize(fyne.NewSize(680, 420))

	d := widget.NewSeparator()
	control_panel := ControlPanelCreate()
	body_panel := BodyCreate()
	main_lay := container.NewBorder(control_panel, nil, nil, nil, d, body_panel) //(widjets.ControlButton("1"), nil, nil, nil)

	w.SetContent(main_lay)

	w.ShowAndRun()
}
