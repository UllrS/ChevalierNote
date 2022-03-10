package view

import (
	"gateway/models"
	"gateway/pkg/control"

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
	body_panel := BodyInit()
	models.BottomLabel = widget.NewLabel("")

	main_lay := container.NewBorder(top_panel, models.BottomLabel, nil, nil, d, body_panel) //(widjets.ControlButton("1"), nil, nil, nil)

	models.AppWindow.SetContent(main_lay)

	if models.FileName != "" {
		control.OpenFileFunc(models.FileName, nil)
	}

	models.AppWindow.ShowAndRun()
}

func BodyInit() *container.Split {
	models.Area = widget.NewMultiLineEntry()
	tree_controlPanel := TreeControlPanel()
	models.TreeView = CreateTree()
	models.List_panel = container.NewBorder(nil, tree_controlPanel, nil, nil, models.TreeView)
	scroll_panel := container.NewVScroll(models.List_panel)

	body_panel := container.NewHSplit(scroll_panel, models.Area)
	body_panel.SetOffset(0.25)
	return body_panel

}
func TreeControlPanel() *widget.Toolbar {
	control_panel := widget.NewToolbar()
	img_minus, _ := fyne.LoadResourceFromPath("assets/image/minus.png")
	img_rename, _ := fyne.LoadResourceFromPath("assets/image/edit.png")
	img_plus, _ := fyne.LoadResourceFromPath("assets/image/plus.png")
	item_newItem := widget.NewToolbarAction(img_plus, control.NewItemDialog)
	item_reName := widget.NewToolbarAction(img_rename, control.RenameItemDialog)
	item_delItem := widget.NewToolbarAction(img_minus, control.DeleteItemDialog)
	control_panel.Append(item_newItem)
	control_panel.Append(item_reName)
	control_panel.Append(item_delItem)
	return control_panel

}
