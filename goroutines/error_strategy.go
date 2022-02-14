package goroutines

import (
	"errors"
	"fmt"
)

func InitErrorStrategy() {

	_, err := f1(-5) //if the value is positive it will not give error or if its negative it will print error i.e (-5 or 5)
	if err != nil {
		fmt.Println("Error =", err) //this can be used instead of panic
		//	panic(err) //its used to give high level security or
	}
	//fmt.Println("Value =", value) // wirite value instead of _ in the line 10 while using this line
}

func f1(arg int) (int, error) {
	if arg < 0 {
		return 0,
			errors.New(" Cant work with negative numbers")

	}
	return arg + 1, nil // its nil coz its right
}
