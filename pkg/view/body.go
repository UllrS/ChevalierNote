package view

import (
	"fmt"
	"gateway/pkg/docmanager"
	"gateway/pkg/logic"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Area *widget.Entry
var list_panel *fyne.Container
var TreeView *widget.Tree

func BodyCreate() *container.Split {
	//b := widget.NewButton("hi!z", Sel)
	Area = widget.NewMultiLineEntry()
	list_panel = container.NewMax()
	scroll_panel := container.NewVScroll(list_panel)

	TreeView = CreateTree()

	list_panel.Add(TreeView)
	body_panel := container.NewHSplit(scroll_panel, Area)
	body_panel.SetOffset(0.25)
	return body_panel

}
func Sel() {
	a := widget.NewButton("New btn", Sel)
	a.Resize(fyne.NewSize(50, 50))
	list_panel.Add(a)
}
func CreateTree() *widget.Tree {

	// a := fyne.CurrentApp()
	tree := &widget.Tree{
		ChildUIDs: func(uid string) []string {
			ess := docmanager.Path_Map[uid]
			list := ess.GetChildrenListName()
			if len(list) < 1 {
				return []string{}
			}
			fmt.Println(list)
			return list
		},
		IsBranch: func(uid string) bool {
			parent, ok := docmanager.Path_Map[uid]
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
			if dataStruct, ok := docmanager.Path_Map[uid]; ok {
				Area.SetText(dataStruct.Note)
				fmt.Println("TREE SELECTED")
				fmt.Println(uid)
				fmt.Println(dataStruct.Note)
			} else {
				fmt.Println("TREE SELECTED ERROR")
			}
		},
		OnUnselected: func(uid string) {
			fmt.Println("TREE UN SELECTED")
			logic.SaveChanges()
			if _, ok := docmanager.Path_Map[uid]; ok {
				docmanager.Path_Map[uid].Note = Area.Text
			} else {
				fmt.Println("TREE UN SELECTED ERROR")
			}
		},
	}
	return tree
}
