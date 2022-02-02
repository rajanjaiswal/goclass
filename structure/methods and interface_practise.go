package structure

import (
	"fmt"
	"math"
)

type geometry interface { // used this as part of interface so that we dont have to type the below lines from line 52 to 61
	area() float64
	perm() float64
}

//method note in using methods and structure pointer and address is converted bt itself

type rect struct {
	width  float64
	height float64
}
type circle struct {
	radius float64
}

//func area(w, h int) int { //// this is old way by not using methods
//return w * h //// this is old way by not using methods
func (r rect) area() float64 { // new way by using methods now
	return r.width * r.height

}

func (r rect) perm() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius

}

func (c circle) perm() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) { // used this as part of interface so that we dont have to type the below lines from line 52 to 61
	fmt.Println("area =", g.area())
	fmt.Println("area =", g.perm())
}

func Init2() {
	r := rect{width: 10, height: 5}
	c := circle{radius: 5}
	//	fmt.Println("Area of rect = ", area(r.height, r.width)) // this is old way by not using methods
	fmt.Println("Area of rect = ", r.area())
	//	r.area()  old way
	fmt.Println("Perimeter of rect = ", r.perm())

	r2 := &r // this for address and pointer
	fmt.Println("Area of rect2 = ", r2.area())
	fmt.Println("Perimeter of rect2 = ", r2.perm())

	fmt.Println("Area of circle = ", c.area())
	fmt.Println("Perimeter of circle = ", c.perm())

	fmt.Println("rect") // used this as part of interface so that we dont have to type the above from line 52 to 61
	measure(r)
	fmt.Println("circle")
	measure(c)

}

//type des interface {
//	describe() string
