package control

import (
	"fmt"
	"gateway/models"
	"gateway/tools"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
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
		fmt.Println(os.Args)
		fmt.Println(os.Args[1])
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
		fmt.Println(uc)
		defer uc.Close()
		if !tools.FileValidator(uc.URI().Extension()) {
			fmt.Println("Error file")
			return
		} else {
			OpenFileFunc(uc.URI().Path(), e)
		}
	}, models.AppWindow)
	filePick.Show()
}
func OpenFileFunc(fpath string, err error) {
	fmt.Println("URI!")
	fmt.Println(fpath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	models.FileName = fpath
	DataLoad()
}
