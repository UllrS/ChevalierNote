package view

import (
	"gateway/pkg/control"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ControlPanelCreate() *widget.Toolbar {

	control_panel := widget.NewToolbar()
	// init resources
	img_save, _ := fyne.LoadResourceFromPath("assets/image/save.png")
	img_saveAs, _ := fyne.LoadResourceFromPath("assets/image/saveas.png")
	img_shield, _ := fyne.LoadResourceFromPath("assets/image/lock.png")
	img_earth, _ := fyne.LoadResourceFromPath("assets/image/earth1.png")
	img_openfolder, _ := fyne.LoadResourceFromPath("assets/image/open_folder.png")
	img_newItem, _ := fyne.LoadResourceFromPath("assets/image/docplus.png")
	//init toolbar button
	item_newItem := widget.NewToolbarAction(img_newItem, control.NewFile)
	item_openfolder := widget.NewToolbarAction(img_openfolder, control.OpenFile)
	item_save := widget.NewToolbarAction(img_save, control.SaveFile)
	item_saveAs := widget.NewToolbarAction(img_saveAs, control.SaveFileAs)
	item_web := widget.NewToolbarAction(img_earth, control.WebS)
	item_guard := widget.NewToolbarAction(img_shield, control.Guard)

	control_panel.Append(item_newItem)
	control_panel.Append(item_openfolder)
	control_panel.Append(item_save)
	control_panel.Append(item_saveAs)
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(item_web)
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(item_guard)
	return control_panel
}
