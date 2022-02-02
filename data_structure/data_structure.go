package datastructure

import (
	"fmt"
)

func Init() {
	practiseMap()
	practiseRange()
	practiseVariadicFunction()

}
func practiseMap() {
	m := make(map[string]int) //map[key-type]value-type // there we can also use string instead of int

	m["a"] = 98
	m["b"] = 97
	m["a"] = 99

	// key is always unique i.e if "a" has two value it will take the latest value

	delete(m, "b") // will delete "b"

	//value, ok := m["a"]; ok can be used here but we used down coz in go we can used it in inline as down down
	fmt.Println("length =", len(m))

	if value, ok := m["a"]; ok { //ok = true if key exist in map i.e 97 if not but default will be 0g
		fmt.Println(ok)
		fmt.Println(value)
	}

	m2 := map[string]string{"a": "kathmandu", "b": "pokhara"} //  done direct inline but have to use make command as above
	fmt.Println("m2", m2)

}
func practiseRange() {
	arr1 := []int{1, 2, 3, 4, 5, 6}
	//for i := 0; index < len(arr1); index++ { // instead of this we can use range check down
	//	fmt.Println(arr1[index])
	//	}
	for _, value := range arr1 { //if i need index i will wright index insted of _ and print it but cant omit index we have to used _
		//fmt.Println(index)
		fmt.Println(value)

	}
	m2 := map[string]string{"a": "kathmandu", "b": "pokhara"} // inline initialization of map

	for k, v := range m2 { // in map it always key and value i.e k and v // iterates over the map m2
		fmt.Println("key =", k)
		fmt.Println("value =", v)

	}

}

func practiseVariadicFunction() { // n no of arguments of type int
	nums := []int{6, 2, 1, 4, 8, 3}
	sum(6, 2, 3, 4, 5)
	sum(1, 2)
	sum(1, 2, 3, 4, 7)
	sum(nums...)

}
func sum(nums ...int) { // this is variadic function ...int or ...string
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

	//annynomous functions i.e i can use it instance whereever i want to use
	sum := func(nums ...int) {
		total := 0
		for _, num := range nums {
			total += num
		}
		fmt.Println(total)
	}
	sum(1, 2, 3)
	// THE BELOW PROGRAM IS FUNCTION INSIDE THE FUNCTION  I.E CLOSERS
	nextSeq := seq()
	fmt.Println(nextSeq()) // i++ = 1
	fmt.Println(nextSeq()) // i++ = 2
	fmt.Println(nextSeq()) // i++ = 3

	nextSeq1 := seq()
	fmt.Println(nextSeq1()) // i++ = 1
	fmt.Println(nextSeq1()) // i++ = 2

}
func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
