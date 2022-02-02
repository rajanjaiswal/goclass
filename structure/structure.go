package structure

import "fmt"

type person struct {
	name string
	age  uint8 // unit8 is used coz age max of any person an be 120 max so to same the memory
	//address string //it can be used if i only need one info i.e only one address
	address map[string]string // use map if u need more then one address or info see below at address
	married bool
}

func Init() {
	//	m := make(map[string]interface{})
	//	m["age"] = 1
	//	m["name"] = "ram"
	//	m["sddress"] ="abc"
	//	m["married"] = false

	p := person{
		name:    "Ram",
		age:     25,
		address: map[string]string{"home": "abc", "office": "def", "rent": "ghk"},
		married: false,
	}
	fmt.Printf("%+v\n", p)

	p1 := person{
		name: "shyam",

		address: map[string]string{"home": "abc", "office": "def", "rent": "ghk"},
	}

	p1.age = 26
	p1.married = true
	fmt.Println(p1.name)
	fmt.Printf("%+v\n", p1)

	// to use address and pointer below is the program
	p2 := newPerson("Hari", 25)
	fmt.Printf("%+v\n", p2)

}

func newPerson(name string, age uint8) *person {
	p := &person{name: name, age: age}
	return p
}
