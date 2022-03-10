package tools

import (
	"fmt"
	"gateway/models"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func EncryptData(data []byte) ([]byte, error) {
	if !models.FileStructure.Settings.Lock {
		return data, nil
	}
	unlockData, err := DataLock(data, models.PWD)
	return unlockData, err
}
func PwdEntry(data []byte) ([]byte, error) {
	if !models.FileStructure.Settings.Lock {
		fmt.Println("NO LOCK")
		return data, nil
	}
	unlockData, err := DataUnlock(data, models.PWD)
	return unlockData, err
}
func PwdCreate() {
	if models.FileStructure.Settings.Lock {
		var entr = widget.NewEntry()
		var entrNew = widget.NewEntry()
		dialog := dialog.NewForm("Новый пароль", "Ок", "Отмена", []*widget.FormItem{widget.NewFormItem("Пароль", entr), widget.NewFormItem("New password", entrNew)}, func(b bool) {
			if b {
				if models.PWD == entr.Text {
					models.PWD = entrNew.Text
					models.FileStructure.Settings.Lock = true
					Alert("Пароль изменен!", "")
				}
			}
		}, models.AppWindow)
		dialog.Show()
	} else {
		var entr = widget.NewEntry()
		dialog := dialog.NewForm("Новый пароль", "Ок", "Отмена", []*widget.FormItem{widget.NewFormItem("Password", entr)}, func(b bool) {
			if b {
				models.PWD = entr.Text
				models.FileStructure.Settings.Lock = true
				Alert("Пароль установлен!", "")
			}
		}, models.AppWindow)
		dialog.Show()
	}
}
