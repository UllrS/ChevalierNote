package control

import (
	"fmt"
	"gateway/models"
	"gateway/pkg/repository"
	"gateway/tools"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func DataSave(path string) {
	repository.SaveFile(path)
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

func SaveChanges(uid string) {
	fmt.Println(uid)
	if ess, ok := models.Path_Map[uid]; ok {
		ess.ChengeNote(models.Area.Text)
		models.Path_Map[uid].Note = models.Area.Text
	} else {
		tools.Alert("ERROR", "SaveChanges: Ошибка выделенной секции")
	}

}
