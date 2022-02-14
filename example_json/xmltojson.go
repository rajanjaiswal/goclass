package examplejson

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Person1 struct {
	FirstName string
	LastName  string
	Age       int
	Married   bool
	//Address   map[string]string
	//	Address Address // i can use this instead of using map but i have to create different type address as line 23-25 below
	//Phone []map[string]string
}

func Initxmltojson() {

	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>

    <first_name>Ram</first_name>
    <last_name>Berma</last_name>
    <age>24</age>
    <married>false</married>
    <first_name>Ram2</first_name>
    <last_name>Berma2</last_name>
    <age>16</age>
    <married>false</married>
    
	`)

	data := &Person1{}
	err := xml.Unmarshal(xmlData, data)
	if nil != err {
		fmt.Println("Error unmarshalling from XML", err)
		return
	}
	result, err := json.Marshal(data)
	if nil != err {
		fmt.Println("Error marshalling to JSON", err)
		return
	}

	fmt.Printf("%s\n", result)

}
