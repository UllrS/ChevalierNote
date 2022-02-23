package control

import (
	"fmt"
	"gateway/models"
	"gateway/pkg/repository"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func DataSave() {
	models.Path_Map[models.TargetEssens].Note = models.Area.Text
	if models.FileName == "" {
		dial := dialog.NewFileSave(func(uc fyne.URIWriteCloser, e error) { repository.SaveFile(uc.URI().Fragment()) }, models.AppWindow)
		dial.Show()
	}
	repository.SaveFile(models.FileName)
}
func DataLoad() {
	repository.LoadFile(models.FileName)
}
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
	models.Path_Map[models.TargetEssens].Rename(newName)
	models.Path_Map = map[string]*models.Essen{}
	models.Essenx.InitPath("")
	models.TreeView.Select(newName)
	models.TreeView.Refresh()
}
func DeleteItemDialog() {
	dialog := dialog.NewConfirm("Are you sure?", "", func(b bool) {
		if b {
			fmt.Println("DELETE!")
		}
	}, models.AppWindow)
	dialog.Show()

}
func DeleteItem() {
	fmt.Println("delete item")
	models.TreeView.Select("")
	// models.Path_Map[models.TargetEssens]
	models.Path_Map = map[string]*models.Essen{}
	models.Essenx.InitPath("")
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
func OpenFile() {
	filePick := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) { fmt.Sprintln(e.Error()) }, models.AppWindow)
	filePick.Show()
	fmt.Println("NEW safe")
}
