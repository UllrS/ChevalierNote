package control

import (
	"gateway/models"
	"gateway/tools"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
)

func NewFile() {
	dialog := dialog.NewConfirm("Are you sure?", "Create new document", func(b bool) {
		if b {
			NewData()
		}
	}, models.AppWindow)
	dialog.Show()

}
func CheckARGS() bool {
	if len(os.Args) > 1 {
		models.FileName = os.Args[1]
		DataLoad()
		return true
	}
	return false
}
func OpenFile() {
	filePick := dialog.NewFileOpen(func(uc fyne.URIReadCloser, e error) {
		if uc == nil {
			return
		}
		defer uc.Close()
		if !tools.FileValidator(uc.URI().Extension()) {
			tools.Alert("ERROR", "Неверный формат файла")
			return
		} else {
			OpenFileFunc(uc.URI().Path(), e)
		}
	}, models.AppWindow)
	filePick.SetFilter(storage.NewExtensionFileFilter([]string{".cvl"}))
	filePick.Show()
}
func OpenFileFunc(fpath string, err error) {
	if err != nil {
		tools.Alert("ERROR", err.Error())
		return
	}
	models.FileName = fpath
	DataLoad()
}
func SaveFile() {
	SaveChanges(models.TargetEssens)
	if _, err := os.Stat(models.FileName); os.IsNotExist(err) || models.FileName == "" {
		SaveFileAs()
		return
	}
	DataSave(models.FileName)
}

func SaveFileAs() {
	if _, ok := models.Path_Map[models.TargetEssens]; ok {
		models.Path_Map[models.TargetEssens].Note = models.Area.Text
	}
	dial := dialog.NewFileSave(func(uc fyne.URIWriteCloser, e error) {
		if uc == nil || e != nil {
			return
		}
		defer uc.Close()
		// uc.URI().Extension()
		models.FileName = uc.URI().Path()
		DataSave(models.FileName)
	}, models.AppWindow)
	dial.SetFilter(storage.NewExtensionFileFilter([]string{".cvl"}))
	dial.SetFileName("newfile.cvl")
	dial.Show()
}
