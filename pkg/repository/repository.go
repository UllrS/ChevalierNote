package repository

import (
	"encoding/json"
	"fmt"
	"gateway/models"
	"os"
)

func LoadFile(docname string) {
	file, err := os.Open(docname)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&models.Essenx)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(models.Essenx)
	fmt.Println(models.Essenx.GetChildrenListName())
}
func SaveFile(docname string) {
	file, err := os.OpenFile(docname, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0640)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(&models.Essenx)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(models.Essenx)
	models.Essenx.InitPath("")

}
