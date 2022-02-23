package control

import (
	"fmt"
	"gateway/models"
	"gateway/pkg/repository"
	"gateway/tools"
	"strings"

	"fyne.io/fyne/v2/dialog"
)

func ShowBar() {
	repository.LoadFile(models.FileName)
}
func NewItemDialog() {
	dialog := dialog.NewEntryDialog("New name", "", func(s string) { NewItem(s) }, models.AppWindow)
	dialog.Show()
}
func NewItem(name string) {
	if _, ok := models.Path_Map[models.TargetEssens]; !ok {
		alert := dialog.NewInformation("Уже существует", "", models.AppWindow)
		alert.Show()
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
	dialog := dialog.NewEntryDialog("New name", "", func(s string) { RenameItem(s) }, models.AppWindow)
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
			fmt.Println("DELETE!")
			DeleteItem()
		}
	}, models.AppWindow)
	dialog.Show()
}
func DeleteItem() {
	if models.TargetEssens == "" {
		return
	}
	fmt.Println("delete item")
	var path_str = models.Path_Map[models.TargetEssens].Path
	fmt.Println(path_str)
	var arr = strings.Split(path_str, "/")
	fmt.Println(arr)
	var itemName = arr[len(arr)-1]
	fmt.Println(itemName)
	var parentName = arr[len(arr)-2]
	arr = arr[:len(arr)-1]
	fmt.Println(arr)
	path_str = strings.Join(arr, "/")
	fmt.Println(path_str)

	models.TreeView.Select(parentName)
	models.Path_Map[parentName].Children = tools.Remove(models.Path_Map[parentName].Children, models.Path_Map[itemName])
	models.Init()

	// models.Path_Map = map[string]*models.Essen{}
	// models.Essenx.InitPath("")
	models.TreeView.Refresh()
}
func WebS() {
	fmt.Println("NEW webs")
}
func Guard() {
	fmt.Println("NEW guard")
}
func NewSafe() {
	fmt.Println("NEW safe")
}
