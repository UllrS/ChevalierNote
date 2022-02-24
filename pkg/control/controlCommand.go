package control

import (
	"fmt"
	"gateway/models"
	"gateway/pkg/repository"
	"gateway/tools"
	"strings"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowBar() {
	repository.LoadFile(models.FileName)
}
func NewItemDialog() {
	// win := dialog.NewEntryDialog("New name", "", func(s string) { NewItem(s) }, models.AppWindow)
	var entr = widget.NewEntry()
	win := dialog.NewForm("Создать раздел", "Ок", "Отмена", []*widget.FormItem{widget.NewFormItem("", entr)}, func(b bool) {
		if b {
			NewItem(entr.Text)
		}
	}, models.AppWindow)

	win.Show()
}
func NewItem(name string) {
	if _, ok := models.Path_Map[models.TargetEssens]; !ok {
		tools.Alert("Не выбран родительский раздел", "")
		dialog.NewInformation("Уже существует", "", models.AppWindow)
		return
	}
	if _, ok := models.Path_Map[name]; ok {
		tools.Alert("Раздел уже существует", "")
		dialog.NewInformation("Уже существует", "", models.AppWindow)
		return
	}
	models.Path_Map[models.TargetEssens].CreateChild(name)
	models.Path_Map = map[string]*models.Essen{}
	models.Essenx.InitPath("")
	models.TreeView.OpenBranch(models.TargetEssens) //ChildUIDs(models.TargetEssens)
	models.TreeView.Select(name)
	models.TreeView.Refresh()
}
func RenameItemDialog() {
	var entr = widget.NewEntry()
	dialog := dialog.NewForm("Новое название раздела", "Ок", "Отмена", []*widget.FormItem{widget.NewFormItem("ENT", entr)}, func(b bool) {
		if b {
			RenameItem(entr.Text)
		}
	}, models.AppWindow)
	dialog.Show()
}

func RenameItem(newName string) {
	if models.TargetEssens == "" {
		return
	}
	models.Path_Map[models.TargetEssens].Rename(newName)
	//models.Path_Map = map[string]*models.Essen{}
	models.Init()
	models.TreeView.Select(newName)
	models.TreeView.Refresh()
}
func DeleteItemDialog() {
	dialog := dialog.NewConfirm("Are you sure?", "Delete selected section", func(b bool) {
		if b {
			DeleteItem()
		}
	}, models.AppWindow)
	dialog.Show()
}
func DeleteItem() {
	if models.TargetEssens == "" {
		return
	}
	var path_str = models.Path_Map[models.TargetEssens].Path
	var arr = strings.Split(path_str, "/")
	var itemName = arr[len(arr)-1]
	var parentName = arr[len(arr)-2]
	//arr = arr[:len(arr)-1]
	//path_str = strings.Join(arr, "/")

	models.TreeView.Select(parentName)
	models.Path_Map[parentName].Children = tools.Remove(models.Path_Map[parentName].Children, models.Path_Map[itemName])
	models.Init()
	models.TreeView.Refresh()
}
func WebS() {
	fmt.Println("NEW webs")
}
func Guard() {
	fmt.Println("NEW guard")
	tools.PwdCreate()
}
func NewSafe() {
	fmt.Println("NEW safe")
}
