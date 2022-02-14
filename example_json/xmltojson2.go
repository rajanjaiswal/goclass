package examplejson

import (
	"encoding/xml"
	"fmt"
)

type Per struct {
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Age       int               `json:"age"`
	Married   bool              `json:"married"`
	Address   map[string]string `json:"address"`
	//	Address Address // i can use this instead of using map but i have to create different type address as line 23-25 below
	Phone []map[string]string `json:"phone_number"`
}

type Response struct {
	Page int      `json:"page"`
	Data []Person `json:"data"`
}

//type Address struct{
//	Country string
//	Street string
//}

func Initxmltojson2() {
	exJson := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
	<page>1</page>
	<data>
		<first_name>Ram</first_name>
		<last_name>Berma</last_name>
		<age>24</age>
		<married>false</married>
		<address>
			<country>Nepal</country>
			<street>butwal</street>
		</address>
		<phone_number>
			<home>9857687987</home>
		</phone_number>
		<phone_number>
			<office>879869860</office>
		</phone_number>
	</data>
	<data>
		<first_name>Ram</first_name>
		<last_name>Berma</last_name>
		<age>24</age>
		<married>false</married>
		<address>
			<country>Nepal</country>
			<street>butwal</street>
		</address>
		<phone_number>
			<home>9857687987</home>
		</phone_number>
		<phone_number>
			<office>879869860</office>
		</phone_number>
	</data>`) //this is injson so changing into map line 92 and 93

	var ex Response

	//var ex8 map[string]interface{} //this is still in map so its difficult to work with also
	// we have already create the structure above so using line 101
	err := xml.Unmarshal(exJson, &ex)
	if err != nil {
		panic((err))

	}
	fmt.Println(ex)
	fmt.Printf("%+v", ex)
}
