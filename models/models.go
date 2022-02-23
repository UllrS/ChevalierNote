package models

import (
	"fyne.io/fyne/v2/widget"
)

func Init() {

	TreeItemMap = map[string]*widget.Label{}
	Path_Map = map[string]*Essen{"": &Essenx}
	Essenx.InitPath("")
}
