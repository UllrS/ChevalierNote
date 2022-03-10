package models

import (
	"fmt"
)

var FileName string
var Essenx Essen
var FileStructure FileStruct
var TargetEssens string

var Path_Map map[string]*Essen

type FileStruct struct {
	Settings FileSettings
	Body     []byte
}
type FileSettings struct {
	Lock bool
}

type Essen struct {
	Name     string   `json:"name"`
	Note     string   `json:"note"`
	Path     string   `json:"path"`
	IsOpen   bool     `json:"isopen"`
	Xcurs    int      `json:"xcurs"`
	Ycurs    int      `json:"ycurs"`
	Children []*Essen `json:"children"`
}

func (e *Essen) CreateChild(name string) {
	var newEss = Essen{Name: name, Path: e.Path + "/" + name}

	e.Children = append(e.Children, &newEss)
}
func (e *Essen) Rename(name string) {
	e.Name = name
}
func (e *Essen) Delete(name string) {

}
func (e *Essen) InitPath(path string) {
	Path_Map[e.Name] = e
	e.Path = fmt.Sprint(path, "/", e.Name)
	for _, value := range e.Children {
		value.InitPath(e.Path)
	}
}
func (e *Essen) ChengeNote(note string) {
	e.Note = note
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
