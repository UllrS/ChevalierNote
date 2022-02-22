package view

import (
	"gateway/pkg/control"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ControlPanelCreate() *widget.Toolbar {

	control_panel := widget.NewToolbar()
	// init resources
	img_save, _ := fyne.LoadResourceFromPath("assets/image/save.png") //canvas.NewImageFromFile("seets/image/save.png")
	img_shield, _ := fyne.LoadResourceFromPath("assets/image/shield.png")
	img_safe, _ := fyne.LoadResourceFromPath("assets/image/safe.png")
	img_earth, _ := fyne.LoadResourceFromPath("assets/image/earth1.png")
	img_openfolder, _ := fyne.LoadResourceFromPath("assets/image/open_folder.png")
	img_newItem, _ := fyne.LoadResourceFromPath("assets/image/docplus.png")
	img_minus, _ := fyne.LoadResourceFromPath("assets/image/minus.png")
	img_rename, _ := fyne.LoadResourceFromPath("assets/image/rename.png")
	img_show_bar, _ := fyne.LoadResourceFromPath("assets/image/dawnloads.png")
	//init toolbar button
	item_openfolder := widget.NewToolbarAction(img_openfolder, control.DataLoad)
	item_show_bar := widget.NewToolbarAction(img_show_bar, control.ShowBar)
	item_save := widget.NewToolbarAction(img_save, control.DataSave)
	item_newItem := widget.NewToolbarAction(img_newItem, control.NewItem)
	item_minus := widget.NewToolbarAction(img_minus, control.DeleteItem)
	item_rename := widget.NewToolbarAction(img_rename, control.RenameItem)
	item_web := widget.NewToolbarAction(img_earth, control.WebS)
	item_guard := widget.NewToolbarAction(img_shield, control.Guard)
	item_safe := widget.NewToolbarAction(img_safe, control.NewSafe)
	control_panel.Append(item_openfolder)
	control_panel.Append(item_save)
	control_panel.Append(item_newItem)
	control_panel.Append(item_minus)
	control_panel.Append(item_rename)
	// control_panel.Append(widget.NewToolbarSeparator())
	control_panel.Append(item_web)
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(item_guard)
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(item_safe)
	control_panel.Append(item_show_bar)
	return control_panel
}
