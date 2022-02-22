package docmanager

import (
	"encoding/json"
	"fmt"
	"os"
)

var Essenx Essen

// var ListName = []string{}
var Path_Map = map[string]*Essen{"": &Essenx}

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
func (e *Essen) InitPath() {
	Path_Map[e.Name] = e
	for _, value := range e.Children {
		value.InitPath()
	}
}
func (e *Essen) GetChildrenListName() []string {
	children_list := []string{}
	if e.Children == nil {
		return children_list
	}
	for _, value := range e.Children {
		children_list = append(children_list, value.Name)
	}
	return children_list
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
	Essenx.InitPath()
	Essenx.createView()
}
func WriteDoc(docname string) {
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
	Essenx.InitPath()
	Essenx.createView()

}
func WriteTest() {
	//...
}
