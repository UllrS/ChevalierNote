package control

import (
	"fmt"
	"gateway/models"
	"gateway/tools"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
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
	fmt.Println("Check args")
	if len(os.Args) > 1 {
		abs, err := filepath.Abs(os.Args[1])
		if err != nil {
			fmt.Println("Error", err.Error())
		}
		models.FileName = abs
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

	var entr = widget.NewEntry()
	dialog := dialog.NewForm("Введите пароль", "Ок", "Отмена", []*widget.FormItem{widget.NewFormItem("ENT", entr)}, func(b bool) {
		if b {
			models.PWD = entr.Text
			models.FileName = fpath
			DataLoad()

		} else {
			models.FileName = fpath
			DataLoad()
		}
	}, models.AppWindow)
	dialog.Show()

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
		} else {
			models.FileName = uc.URI().Path()
			DataSave(models.FileName)
		}
		defer uc.Close()
		// uc.URI().Extension()
	}, models.AppWindow)
	dial.SetFilter(storage.NewExtensionFileFilter([]string{".cvl"}))
	dial.SetFileName("newfile.cvl")
	dial.Show()
}
