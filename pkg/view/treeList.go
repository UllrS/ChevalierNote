package view

import (
	"gateway/models"
	"gateway/pkg/control"
	"gateway/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CreateTree() *widget.Tree {

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
			// var btn_rename = widget.NewButton("rename", control.RenameItemDialog)
			// var btn_del = widget.NewButton("del", control.DeleteItemDialog)
			var lbl = widget.NewLabel("")
			return lbl
		},
		UpdateNode: func(uid string, branch bool, obj fyne.CanvasObject) {
			models.TreeItemMap[uid] = obj.(*widget.Label)
			models.TreeItemMap[uid].SetText(uid)
			// obj.(*widget.Label).SetText(uid)
		},
		OnSelected: func(uid string) {
			models.TargetEssens = uid
			if dataStruct, ok := models.Path_Map[uid]; ok {
				models.Area.SetText(dataStruct.Note)
				models.BottomLabel.SetText(dataStruct.Path)
			} else {
				tools.Alert("ERROR", "OnSelected: Ошибка выделенной секции")
			}
		},
		OnUnselected: func(uid string) {
			if _, ok := models.Path_Map[uid]; uid != "" && ok {
				control.SaveChanges(uid)
			}
		},
	}
	return tree
}
