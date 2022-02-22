package widjets

import (
	"fmt"

	"fyne.io/fyne/v2/widget"
)

func ControlButton(btn_name string) *widget.Button {
	btn := widget.NewButton(btn_name, func() { fmt.Println("Control_panel") })
	return btn
}
