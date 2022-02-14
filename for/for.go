package forr

import "fmt"

func Forr() {
	sum := 0
	for j := 0; j < 3; j++ {
		for i := 0; i < 10; i++ {
			if i == 7 {
				break
			}
			if i == 5 {
				continue
			}
			sum += i
			fmt.Println(sum)
		}
	}
	fmt.Println("Total sum = ", sum)
	sum = 1
	for sum < 1000 {
		sum += sum // sum + sum
	}
	fmt.Println("total sum of 1000 =", sum)
	sum += sum

}
