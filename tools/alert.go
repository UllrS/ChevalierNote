package tools

import (
	"gateway/models"

	"fyne.io/fyne/v2/dialog"
)

func Alert(title string, message string) {
	alert := dialog.NewInformation(title, message, models.AppWindow)
	alert.Show()
}
