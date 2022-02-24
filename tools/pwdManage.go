package tools

import (
	"gateway/models"

	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func PwdEntry(data []byte) []byte {
	if data[0] == 01 {
		models.Lock = true
		//f Decode aes256
		//
	} else if data[0] == 00 {
		models.Lock = false
	}
	data = data[1:]
	if !models.Lock {
		return data
	}
	var unlockData []byte
	var entr = widget.NewEntry()
	dialog := dialog.NewForm("Введите пароль", "Ок", "Отмена", []*widget.FormItem{widget.NewFormItem("ENT", entr)}, func(b bool) {
		if b {
			unlockData = DataUnlock(data, entr.Text)
			models.PWD = entr.Text
		}
	}, models.AppWindow)
	dialog.Show()
	return unlockData
}
func PwdCreate() {
	if models.Lock {
		Alert("ERROR", "Файл уже зашифрован")
		return
	}
	var entr = widget.NewEntry()
	dialog := dialog.NewForm("Новый пароль пароль", "Ок", "Отмена", []*widget.FormItem{widget.NewFormItem("ENT", entr)}, func(b bool) {
		if b {
			models.PWD = entr.Text
			models.Lock = true
			Alert("Пароль установлен!", "")
		}
	}, models.AppWindow)
	dialog.Show()
}
