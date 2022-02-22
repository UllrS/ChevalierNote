package view

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ControlPanelCreate() *widget.Toolbar {

	control_panel := widget.NewToolbar()
	img_save, _ := fyne.LoadResourceFromPath("assets/image/save.png") //canvas.NewImageFromFile("seets/image/save.png")
	img_shield, _ := fyne.LoadResourceFromPath("assets/image/shield.png")
	img_safe, _ := fyne.LoadResourceFromPath("assets/image/safe.png")
	img_earth, _ := fyne.LoadResourceFromPath("assets/image/earth1.png")
	item := widget.NewToolbarAction(img_save, func() { fmt.Println("Control_panel") })
	item2 := widget.NewToolbarAction(img_earth, func() { fmt.Println("Control_panel") })
	item3 := widget.NewToolbarAction(img_shield, func() { fmt.Println("Control_panel") })
	item4 := widget.NewToolbarAction(img_safe, func() { fmt.Println("Control_panel") })
	control_panel.Append(item)
	control_panel.Append(widget.NewToolbarSeparator())
	control_panel.Append(item2)
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(item3)
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(item4)
	return control_panel
}
