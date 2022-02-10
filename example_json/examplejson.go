package examplejson

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string            `json:"first_name"`
	LastName  string            `json:"last_name"`
	Age       int               `json:"age"`
	Married   bool              `json:"married"`
	Address   map[string]string `json:"address"`
	//	Address Address // i can use this instead of using map but i have to create different type address as line 23-25 below
	Phone []map[string]string `json:"phone_number"`
}

type response struct {
	Page int      `json:"page"`
	Data []Person `json:"data"`
}

//type Address struct{
//	Country string
//	Street string
//}

func InitExampleJson() {

	// key, value
	//Json Object its signal data
	//Json Array its multipal data
	// data type: number, string, bool
	ex1, _ := json.Marshal(true)
	fmt.Println(string(ex1))

	ex2, _ := json.Marshal(1)
	fmt.Println(string(ex2))

	ex3, _ := json.Marshal("golang")
	fmt.Println(string(ex3))

	exD := []string{"a", "b", "c"}
	ex4, _ := json.Marshal(exD)
	fmt.Println(string(ex4))

	exD2 := map[string]interface{}{"first_name": "Ram", "last_name": "Berma", "age": 24} //key is string and value is interface which can be anything
	fmt.Println(exD2)                                                                    // this is map or regular format
	ex5, _ := json.Marshal(exD2)                                                         //THIS IS IN json format converted from map
	fmt.Println(string(ex5))

	person1 := &Person{ //linked to line 86
		FirstName: "Ram",
		LastName:  "Berma",
		Age:       24,
		Married:   false,
		Address: map[string]string{
			"country": "Nepal",
			"street":  "butwal",
		},
		Phone: []map[string]string{
			{"home": "9857687987"},
			{"office": "879869860"},
		},
	}
	person2 := &Person{
		FirstName: "Ram",
		LastName:  "Berma",
		Age:       24,
		Married:   false,
		Address: map[string]string{
			"country": "Nepal",
			"street":  "butwal",
		},
		Phone: []map[string]string{
			{"home": "9857687987"},
			{"office": "879869860"},
		},
	}
	data := []Person{*person1, *person2} // linked to line 90

	resp1 := &response{
		Page: 1,
		Data: data, //same thing done as line 13 and 14
	}
	fmt.Printf("resp1=%+v\n", resp1) //this is not is json
	ex6, _ := json.Marshal(resp1)    //this is in json
	fmt.Println(string(ex6))         //this is in json

	exJson := []byte(`{"a":1,"b":"abc","c":[10,11]}`) //this is injson so changing into map line 92 and 93

	var ex7 map[string]interface{}
	err := json.Unmarshal(exJson, &ex7)
	if err != nil {
		panic((err))
	}
	fmt.Println(ex7)

	ex2Json := []byte(`{"page":1,
	"data":[{"first_name":"Ram",
	"last_name":"Berma",
	"age":24,
	"married":false,
	"address":{"country":"Nepal","street":"butwal"},
	"phone_number":[{"home":"9857687987"},{"office":"879869860"}]},
	
	{"first_name":"Ram",
	"last_name":"Berma",
	"age":16,
	"married":false,
	"address":{"country":"Nepal","street":"butwal"},
	"phone_number":[{"home":"9857687987"},{"office":"879869860"}]}]}`) //this is injson so changing into map line 92 and 93

	fmt.Println("=================================")
	fmt.Println("Unmarshalling the persons record")

	var ex8 response

	//var ex8 map[string]interface{} //this is still in map so its difficult to work with also
	// we have already create the structure above so using line 101
	err = json.Unmarshal(ex2Json, &ex8)
	if err != nil {
		panic((err))
	}
	fmt.Println(ex8)
	fmt.Printf("%+v\n", ex8)
	fmt.Println(ex8.Page) // this will give me page1 there can be multipal page number
	fmt.Println(ex8.Data) // this will give me data form inside the specific page

	for _, value := range ex8.Data {
		if value.Age < 16 {
			fmt.Printf("%15s %15s\n", value.FirstName, value.LastName) // this will give me first name and last name from inside the data aslo
			// can get address or age or phn

		}
	}

}
