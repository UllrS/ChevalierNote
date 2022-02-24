package repository

import (
	"encoding/json"
	"gateway/models"
	"gateway/tools"
	"os"
)

func LoadFileJ(docname string) {
	file, err := os.Open(docname)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&models.Essenx)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
}
func SaveFileJ(docname string) {
	file, err := os.OpenFile(docname, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0640)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(&models.Essenx)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}

	models.Essenx.InitPath("")

}
