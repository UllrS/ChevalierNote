package view

import (
	"gateway/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func MainWindowInit() {

	icon, _ := fyne.LoadResourceFromPath("assets/image/knight.png")
	a := app.New()
	models.AppWindow = a.NewWindow("Note")
	// models.AppWin = models.AppWindow{&a}
	models.AppWindow.SetIcon(icon)
	models.AppWindow.Resize(fyne.NewSize(680, 420))

	d := widget.NewSeparator()
	control_panel := ControlPanelCreate()
	subControlPanel := container.NewHBox()
	top_panel := container.NewVBox(control_panel, subControlPanel)
	body_panel := BodyCreate()
	models.BottomLabel = widget.NewLabel("")

	main_lay := container.NewBorder(top_panel, models.BottomLabel, nil, nil, d, body_panel) //(widjets.ControlButton("1"), nil, nil, nil)

	models.AppWindow.SetContent(main_lay)

	models.AppWindow.ShowAndRun()
}
