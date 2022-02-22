package logic

import (
	"fmt"
	"gateway/pkg/docmanager"
)

func SaveChanges() {

}
func SaveFile() {
	fmt.Println("item_save")
	docmanager.WriteDoc("fff.json")
}
