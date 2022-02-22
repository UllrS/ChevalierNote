package control

import (
	"fmt"
	"gateway/models"
)

func SaveChanges(uid string) {
	if ess, ok := models.Path_Map[uid]; ok {
		ess.ChengeNote(models.Area.Text)
		models.Path_Map[uid].Note = models.Area.Text
	} else {
		fmt.Println("TREE UN SELECTED ERROR")
	}

}
