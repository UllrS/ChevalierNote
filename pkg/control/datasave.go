package control

import (
	"fmt"
	"gateway/models"
	"gateway/pkg/repository"
)

func DataSave() {
	models.Path_Map[models.TargetEssens].Note = models.Area.Text
	repository.SaveFile(models.FileName)
}
func DataLoad() {
	repository.LoadFile(models.FileName)
}
func ShowBar() {
	repository.LoadFile(models.FileName)
}
func NewItem() {
	models.Path_Map[models.TargetEssens].CreateChild("New")
	models.Path_Map = map[string]*models.Essen{}
	models.Essenx.InitPath()
	models.TreeView.OpenBranch(models.TargetEssens) //ChildUIDs(models.TargetEssens)
	models.TreeView.Select("New")
	models.TreeView.Refresh()
}
func RenameItem() {
	models.Path_Map[models.TargetEssens].Rename("32New")
	models.Path_Map = map[string]*models.Essen{}
	models.Essenx.InitPath()
	models.TreeView.Select("32New")
	models.TreeView.Refresh()
}
func DeleteItem() {
	fmt.Println("delete item")
	models.TreeView.Select("")
	models.Path_Map[models.TargetEssens]
	models.Path_Map = map[string]*models.Essen{}
	models.Essenx.InitPath()
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
