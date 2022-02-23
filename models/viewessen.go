package models

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var AppWindow fyne.Window

var Area *widget.Entry
var List_panel *fyne.Container
var TreeView *widget.Tree
var BottomLabel *widget.Label

var TreeItemMap map[string]*widget.Label

type TreeItem struct {
	Lbl *widget.Label
	Btn *widget.Button
}
