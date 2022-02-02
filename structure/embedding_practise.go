package structure

import (
	"fmt"
)

type des interface {
	describe() string
	subscribe() string
}

type base struct { // this is structure
	des
	num int
}

func (b base) describe() string { // this is method
	return fmt.Sprintf("base num =%v\n", b.num)
}

func (b base) subscribe() string { // this is method
	return fmt.Sprintf("subscribing %s\n", "")
}

type container struct { // embedding base in new structure
	base
	num int
	str string
}

func Init3() {
	con := container{
		base: base{
			num: 10,
		},
		str: "test",
	}
	fmt.Printf("container =%+v\n", con)

	fmt.Println(con.num, con.str)
	fmt.Println(con.base.num, con.str)

	fmt.Println(con.describe())
	fmt.Println(con.subscribe())
}
