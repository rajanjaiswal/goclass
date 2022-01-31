package initials

import (
	"fmt"
)

func Init() {

	const PI = 3.14 //constant deceleration also not decelare by dynamic type ie ":="
	a := byte('A')
	b := 'a'                  //dynamic deceleration aslo it cant be kept outside the func main but var can be kept outside func main
	var x int = -255          //int
	var y uint16 = 255        //uint
	var z float32 = 255.56    //float
	var m, n, d = 3, 4, "foo" //mixed deceleration
	r := 10 / 5
	var num1 int = 32
	var num2 int = 32
	answer := num1 + num2

	fmt.Println(a)
	fmt.Println(b, n, m, d)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(PI)
	fmt.Printf("%c= %d %c = %U\n", a, a, b, b)
	fmt.Printf("%T\n", b)
	fmt.Println("Value of r = ", r)
	fmt.Println("14 == 14", 14 == 14)

	fmt.Printf("%d", answer)
	s := 2
	Add(r, s)
}
func Add(r int, s int) {
	b := r + s

	fmt.Println(b)
}
