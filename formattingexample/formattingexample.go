package formattingexample

import "fmt"

type point struct {
	x, y int
}

func InitFormattingExamples() {
	p := point{10, 20}
	fmt.Printf("format 1 : %v \n", p)

	fmt.Printf("format 2 : %+v \n", p)

	fmt.Printf("format 3 : %#v \n", p)

	fmt.Printf("Type : %T \n", p)

	fmt.Printf("Integer : %d \n", 15)

	fmt.Printf("boolean : %t \n", true)

	fmt.Printf("Binary : %b \n", 8) // will give output in binary number

	fmt.Printf("character : %c \n", 98) // will print the asci value

	fmt.Printf("pointer : %p \n", &p) // for pointer

	fmt.Printf("hex : %x \n", 500)

	fmt.Printf("float : %.1f \n", 12.778) // for float also if i put .1f it will only print one decimel

	fmt.Printf("string : %s \n", "apple")

	fmt.Printf("format 4 : |%4d|%10d| \n", 10, 20)
	fmt.Printf("format 4 : |%-4d|%-10d| \n", 10, 20)

	fmt.Printf("format 5 : |%-10.2f|%10.2f| \n", 10.0, 20.778)

	fmt.Printf("|%15s|%15s| \n", "---------------", "---------------")
	fmt.Printf("|%15s|%15s| \n", "FirstName", "LastName")
	fmt.Printf("|%15s|%15s| \n", "---------------", "---------------")
	fmt.Printf("|%15s|%15s| \n", "Ram", "Sharma")
	fmt.Printf("|%15s|%15s| \n", "Shyam", "Berma")
	fmt.Printf("|%15s|%15s| \n", "Ram", "Sharma")
	fmt.Printf("|%15s|%15s| \n", "Shyam", "Berma")
	fmt.Printf("|%15s|%15s| \n", "Ram", "Sharma")
	fmt.Printf("|%15s|%15s| \n", "Ram", "Sharma")
	fmt.Printf("|%15s|%15s| \n", "Shyam", "Berma")
	fmt.Printf("|%15s|%15s| \n", "Shyam", "Berma")
	fmt.Printf("|%15s|%15s| \n", "---------------", "---------------")
}
