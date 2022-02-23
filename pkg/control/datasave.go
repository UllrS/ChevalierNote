package control

import (
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
func DataSaveAs() {
	models.Path_Map[models.TargetEssens].Note = models.Area.Text
	if models.FileName == "" {
		dial := dialog.NewFileSave(func(uc fyne.URIWriteCloser, e error) { repository.SaveFile(uc.URI().Fragment()) }, models.AppWindow)
		dial.Show()
	}
	repository.SaveFile(models.FileName)
}
func DataLoad() {
	models.Essenx = models.Essen{}
	repository.LoadFile(models.FileName)
	models.Init()
	if models.TreeView != nil {
		models.TreeView.Refresh()
	}
}
func NewData() {
	models.FileName = ""
	models.Area.SetText("")
	models.Essenx = models.Essen{}
	models.Init()
	if models.TreeView != nil {
		models.TreeView.Refresh()
	}
}
