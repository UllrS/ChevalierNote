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
		var entrNew1 = widget.NewEntry()
		dialog := dialog.NewForm("Изменить пароль", "Ок", "Отмена", []*widget.FormItem{widget.NewFormItem("pwd", entr), widget.NewFormItem("New", entrNew), widget.NewFormItem("New", entrNew1)}, func(b bool) {
			if b {
				if entr.Text != models.PWD {
					Alert("Неверный пароль", "")
				}
				if entrNew.Text != entrNew1.Text {
					Alert("Пароли не совпадают", "")
				}
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
		var entr1 = widget.NewEntry()
		dialog := dialog.NewForm("Зашифровать данные", "Ок", "Отмена", []*widget.FormItem{widget.NewFormItem("New", entr), widget.NewFormItem("", entr1)}, func(b bool) {
			if b {
				models.PWD = entr.Text
				models.FileStructure.Settings.Lock = true
				Alert("Пароль установлен!", "")
			}
		}, models.AppWindow)
		dialog.Show()
	}
}
