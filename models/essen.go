package models

type Essen struct {
	Name   string
	Note   string
	Chield []Essen
}
