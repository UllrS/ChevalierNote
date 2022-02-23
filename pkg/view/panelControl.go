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
	img_shield, _ := fyne.LoadResourceFromPath("assets/image/shield.png")
	img_safe, _ := fyne.LoadResourceFromPath("assets/image/safe.png")
	img_earth, _ := fyne.LoadResourceFromPath("assets/image/earth1.png")
	img_openfolder, _ := fyne.LoadResourceFromPath("assets/image/open_folder.png")
	img_newItem, _ := fyne.LoadResourceFromPath("assets/image/docplus.png")
	img_control, _ := fyne.LoadResourceFromPath("assets/image/control.png")
	//init toolbar button
	item_newItem := widget.NewToolbarAction(img_newItem, control.NewFile)
	item_openfolder := widget.NewToolbarAction(img_openfolder, control.OpenFile)
	item_save := widget.NewToolbarAction(img_save, control.DataSave)
	item_saveAs := widget.NewToolbarAction(img_saveAs, control.DataSaveAs)
	item_web := widget.NewToolbarAction(img_earth, control.WebS)
	// item_minus := widget.NewToolbarAction(img_minus, control.DeleteItem)
	// item_rename := widget.NewToolbarAction(img_rename, control.RenameItemDialog)
	item_safe := widget.NewToolbarAction(img_safe, control.NewSafe)
	item_guard := widget.NewToolbarAction(img_shield, control.Guard)
	item_show_bar := widget.NewToolbarAction(img_control, control.ShowBar)
	control_panel.Append(item_newItem)
	control_panel.Append(item_openfolder)
	control_panel.Append(item_save)
	control_panel.Append(item_saveAs)
	// control_panel.Append(item_minus)
	// control_panel.Append(item_rename)
	// control_panel.Append(widget.NewToolbarSeparator())
	control_panel.Append(item_web)
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(widget.NewToolbarSpacer())
	control_panel.Append(item_safe)
	control_panel.Append(item_guard)
	control_panel.Append(item_show_bar)
	return control_panel
}
