package tools

import (
	"fmt"
	"os"
)

func asd() {
	err := os.MkdirAll("logs", 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("asd")
	file, err := os.OpenFile("dd.cvl", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0640)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()
	// file.Write(b)
	// file.Read(b)
}
