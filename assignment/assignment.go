package assignment

import "fmt"

func Assgnment() {
	even := 0
	odd := 0
	for i := 1; i <= 50; i++ {
		if i%2 == 0 {
			even += i
			//	fmt.Println("Total Sum of even: ", even) if used inside it will print even of all the loop
		} else {
			odd += i
			//	fmt.Println("Total Sum of odd: ", odd) if used inside it will print odd of all the loop
			// so we print outside the if else loop
		}

	}
	fmt.Println("Total Sum of even: ", even)
	fmt.Println("Total Sum of odd: ", odd)
}
