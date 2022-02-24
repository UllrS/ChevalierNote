package repository

import (
	"encoding/json"
	"fmt"
	"gateway/models"
	"gateway/tools"
	"io/ioutil"
)

func LoadFile(docname string) {
	fmt.Printf("load file")
	f, err := ioutil.ReadFile(docname)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
	if len(f) > 0 {
		f = tools.PwdEntry(f)
	}
	models.Essenx = models.Essen{}
	err = json.Unmarshal(f, &models.Essenx)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
}
func SaveFile(docname string) {
	//Init checkByte aes256
	var crypto []byte
	if models.Lock {
		crypto = append(crypto, 01)
	} else {
		crypto = append(crypto, 00)
	}
	var jsonData []byte
	jsonData, err := json.Marshal(models.Essenx)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
	//jsonData encode aes256
	//
	//Add checkByte
	jsonData = append(crypto, jsonData...)
	fmt.Println("jsonData")
	fmt.Println(jsonData)
	err = ioutil.WriteFile(docname, jsonData, 0640)
	if err != nil {
		tools.Alert("ERROR", err.Error())
	}
	models.Essenx.InitPath("")

}
