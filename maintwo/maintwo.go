package maintwo

import (
	"fmt"
	"strings"
)

func swap(x, y string) (string, string) {
	return y, x
}

func Two() {

	//a := 2
	//b := 5
	//add(a, b)
	var greeting = "Hello World"
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Count(greeting, "l"))
	//func add(a int, b int) {
	//c := a + b
	//c -= a

	//fmt.Println(c)
	x, y := swap("x", "y")
	fmt.Println(x, y)

}
