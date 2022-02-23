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
			// var btn_del = widget.NewButton("12", nil)
			control_panel := widget.NewToolbar()
			img_minus, _ := fyne.LoadResourceFromPath("assets/image/minus.png")
			item_minus := widget.NewToolbarAction(img_minus, control.DeleteItemDialog)
			img_rename, _ := fyne.LoadResourceFromPath("assets/image/rename.png")
			item_rename := widget.NewToolbarAction(img_rename, control.RenameItemDialog)
			control_panel.Append(item_rename)
			control_panel.Append(item_minus)
			// var btn_rename = widget.NewButton("rename", control.RenameItemDialog)
			// var btn_del = widget.NewButton("del", control.DeleteItemDialog)
			var lbl = widget.NewLabel("")
			var wid = container.NewBorder(nil, nil, nil, control_panel, lbl)
			return wid
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			var d = models.TreeItem{Lbl: obj.(*fyne.Container).Objects[0].(*widget.Label)}
			models.TreeItemMap[uid] = &d
			models.TreeItemMap[uid].Lbl.SetText(uid)
			// for _, val := range obj.(*fyne.Container).Objects {
			// 	if reflect.TypeOf(val) == reflect.TypeOf(&widget.Label{}) {
			// 		val.(*widget.Label).SetText(uid)
			// 	}
			// }
			//obj.(*widget.Label).SetText(uid)
		},
		OnSelected: func(uid string) {
			models.TargetEssens = uid
			if dataStruct, ok := models.Path_Map[uid]; ok {
				models.Area.SetText(dataStruct.Note)
				models.BottomLabel.SetText(dataStruct.Path)
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
