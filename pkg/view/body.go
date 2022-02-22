package view

import (
	"fmt"
	docmanager "gateway/pkg/doc_manager"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var list_panel *fyne.Container

func BodyCreate() *container.Split {
	//b := widget.NewButton("hi!z", Sel)
	area := widget.NewMultiLineEntry()
	list_panel = container.NewMax()
	scroll_panel := container.NewVScroll(list_panel)
	docmanager.Essenx.GetListName()

	list_panel_pr := widget.NewTree(
		func(tni widget.TreeNodeID) []widget.TreeNodeID {
			fmt.Println("List name", docmanager.ListName)
			return docmanager.ListName
		},
		func(tni widget.TreeNodeID) bool {
			fmt.Println("TNI name")
			if tni != "second" {
				return true
			} else {
				return false
			}
		},
		func(b bool) fyne.CanvasObject {
			return widget.NewLabel("NEW label") // widjets.ControlButton("name btn")
		},
		func(tni widget.TreeNodeID, b bool, co fyne.CanvasObject) {
			co.(*widget.Label).Text = tni
		})
	list_panel_pr.Refresh()
	list_panel.Add(list_panel_pr)
	list_panel_pr.Refresh()
	body_panel := container.NewHSplit(scroll_panel, area)
	body_panel.SetOffset(0.25)
	return body_panel

}
func Sel() {
	a := widget.NewButton("New btn", Sel)
	a.Resize(fyne.NewSize(50, 50))
	list_panel.Add(a)
}
