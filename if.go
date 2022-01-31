package iffit

import "fmt"

func Iffit() {
	x := 5 + 1
	//	if x := 5 + 3; x > 9 {
	//		fmt.Println(x)
	//	} else if x > 5 {
	//		fmt.Println("else on if", x)
	//	} else if x > 7 {
	//		fmt.Println("Print on", x)

	//	} else {
	//		fmt.Println(" on", x)
	//	}
	switch x {
	case 9:
		fmt.Println("on", x)
	case 7:
		fmt.Println("it", x)
	case 5:
		fmt.Println("by", x)
	default:
		fmt.Println(x)
	}
}
