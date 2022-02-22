package view

import (
	"fmt"
	"gateway/models"
	"gateway/pkg/control"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func BodyCreate() *container.Split {
	//b := widget.NewButton("hi!z", Sel)
	models.NameEntry = widget.NewEntry()
	models.Area = widget.NewMultiLineEntry()
	models.List_panel = container.NewMax()
	scroll_panel := container.NewVScroll(models.List_panel)

	models.TreeView = CreateTree()

	models.List_panel.Add(models.TreeView)
	body_panel := container.NewHSplit(scroll_panel, models.Area)
	body_panel.SetOffset(0.25)
	return body_panel

}
func CreateTree() *widget.Tree {

	// a := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			ess := models.Path_Map[uid]
			list := ess.GetChildrenListName()
			if len(list) < 1 {
				return []string{}
			}
			return list
		},
		IsBranch: func(uid string) bool {
			parent, ok := models.Path_Map[uid]
			children := parent.GetChildrenListName()
			return ok && len(children) > 0
		},
		CreateNode: func(branch bool) fyne.CanvasObject {
			return widget.NewLabel("Collection Widgets")
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(uid)
		},
		OnSelected: func(uid string) {
			models.TargetEssens = uid
			if dataStruct, ok := models.Path_Map[uid]; ok {
				models.Area.SetText(dataStruct.Note)
				fmt.Println("TREE SELECTED")
			} else {
				fmt.Println("TREE SELECTED ERROR")
			}
		},
		OnUnselected: func(uid string) {
			fmt.Println("TREE UN SELECTED")
			control.SaveChanges(uid)
		},
	}
	return tree
}
