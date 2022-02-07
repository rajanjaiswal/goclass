package goroutines

import "fmt"

func InitDeferExamples() {
	defer fmt.Println("program finished 1")
	defer fmt.Println("program finished 2")
	value := max(10, 5)
	fmt.Println("max value =", value)
	fmt.Println("Program finished")
}

func max(num1, num2 int) int {
	defer fmt.Println("end of the max function")
	fmt.Println("calculating max value")
	if num1 > num2 {
		return num1

	} else {
		return num2
	}
}
