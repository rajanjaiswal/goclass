package arraypractise

import "fmt"

func Test() {
	//arraypractise()
	slicesPractise()

}
func slicesPractise() {
	s := make([]int, 3, 4) // 2 is length and 6 is max capacity it can hold but we can increase using append
	fmt.Println("slice initial :=", s)
	fmt.Printf("length =%d, capacity =%d \n", len(s), cap(s))
	s[0] = 1
	s[1] = 2
	fmt.Println("slice initial 1 :=", s)

	s = append(s, 4)
	s = append(s, 3)
	fmt.Println("slice append :=", s)
	fmt.Printf("length =%d, capacity =%d \n", len(s), cap(s))

	s1 := make([]int, len(s))
	copy(s1, s)
	fmt.Println("s1 after copy :=", s1)

	l := s[2:4] //s[lower : upper] lower is included; upper is excluded
	fmt.Println("slice second :=", l)

	l = s[:4] //0 to 4-1(3)
	fmt.Println("slice third :=", l)

	l = s[1:] // 1 to last index
	fmt.Println("slice fourth :=", l)

	singlLineSlice := []string{"a", "b", "c"}
	l = s[:4] //0 to 4-1(3)
	fmt.Println("slice fifth :=", singlLineSlice)

	tD := make([][]int, 4)
	for i := 0; i < 4; i++ {
		internaArrayLen := i + 1
		tD[i] = make([]int, internaArrayLen)
		for j := 0; j < internaArrayLen; j++ {
			tD[i][j] = i + j
		}
	}
	fmt.Println("2d slice =", tD)

}

func arraypractise() {
	var a [4]int //[0, 0, 0, 0]  // a[length of array]data_type
	fmt.Println("array :=", a)
	a[2] = 4 //a[index] = value
	fmt.Println("Array :=", a)
	fmt.Println("2nd index :=", a[2])

	fmt.Println("array :=", len(a))

	b := [3]int{1, 2, 3} //giging initial array in a line
	fmt.Println("Initial array :=", b)

	var tD [2][3]int //[[0, 0, 0], [0, 0, 0]]
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			tD[i][j] = i + j //tD[0][0] = 0; tD[0][1] = 1; tD[0][2] = 2
			//tD[1][0] = 1; tD[1][1] = 2; tD[1][2] = 3
		}
	}
	fmt.Println("2nd array =", tD)

}
