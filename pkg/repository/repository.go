package repository

import (
	"encoding/json"
	"fmt"
	"gateway/models"
	"gateway/tools"
	"io/ioutil"
)

func LoadFile(docname string) error {
	fmt.Printf("load file")
	f, err := ioutil.ReadFile(docname)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
	models.FileStructure = models.FileStruct{}

	if err := json.Unmarshal(f, &models.FileStructure); err != nil {
		tools.Alert("ERROR json file structure", err.Error())
		return err
	}

	unlock_body, err := tools.PwdEntry(models.FileStructure.Body)
	if err != nil {
		tools.Alert("ERROR password", err.Error())
		return err
	}
	if err := json.Unmarshal(unlock_body, &models.Essenx); err != nil {
		tools.Alert("ERROR json2", err.Error())
		return err
	}
	return nil
}
func SaveFile(docname string) {
	//Init checkByte aes256
	var jsonData []byte
	jsonData, err := json.Marshal(models.Essenx)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
	// lock data
	jsonData, _ = tools.EncryptData(jsonData)

	// models.FileStructure = models.FileStruct{Body: jsonData}
	models.FileStructure.Body = jsonData
	fileData, err := json.Marshal(models.FileStructure)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
	err = ioutil.WriteFile(docname, fileData, 0640)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
	models.Essenx.InitPath("")

}
