package docmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

var Essenx Essen
var ListName = []string{}

type Essen struct {
	Name     string  `json:"name"`
	Note     string  `json:"note"`
	Children []Essen `json:"children"`
}

func (e *Essen) createView() {
	fmt.Println("createView")
	fmt.Println(e.Name)
	fmt.Println(e.Note)
	for _, value := range e.Children {
		value.createView()
	}
}
func (e *Essen) GetListName() {
	fmt.Println("createView")
	ListName = append(ListName, e.Name)
	fmt.Println(e.Name)
	fmt.Println(e.Note)
	for _, value := range e.Children {
		value.GetListName()
	}

}

func ReadDoc(docname string) {
	file, err := os.Open(docname)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	Essenx = Essen{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Essenx)
	if err != nil {
		fmt.Println(err.Error())
	}
	Essenx.createView()
}
func WriteDoc(docname string) {
	Essenx = Essen{
		Name: "first",
		Note: "first note 123225",
		Children: []Essen{
			{
				Name:     "second",
				Note:     "second note 123225",
				Children: []Essen{},
			},
		},
	}

	file, err := os.OpenFile(docname, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0640)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&Essenx)
	if err != nil {
		fmt.Println(err.Error())
	}
	Essenx.createView()

}
func WriteTest() {
	//...
}
