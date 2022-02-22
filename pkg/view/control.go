package view

import (
	"fmt"
	"gateway/pkg/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ControlPanelCreate() *widget.Toolbar {

	control_panel := widget.NewToolbar()
	img_save, _ := fyne.LoadResourceFromPath("assets/image/save.png") //canvas.NewImageFromFile("seets/image/save.png")
	img_shield, _ := fyne.LoadResourceFromPath("assets/image/shield.png")
	img_safe, _ := fyne.LoadResourceFromPath("assets/image/safe.png")
	img_earth, _ := fyne.LoadResourceFromPath("assets/image/earth1.png")
	item_save := widget.NewToolbarAction(img_save, logic.SaveFile)
	item_web := widget.NewToolbarAction(img_earth, func() { fmt.Println("Control_panel") })
	item_guard := widget.NewToolbarAction(img_shield, func() { fmt.Println("Control_panel") })
	item4 := widget.NewToolbarAction(img_safe, func() { fmt.Println("Control_panel") })
	control_panel.Append(item_save)
	control_panel.Append(widget.NewToolbarSeparator())
	control_panel.Append(item_web)
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(item_guard)
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(item4)
	return control_panel
}
