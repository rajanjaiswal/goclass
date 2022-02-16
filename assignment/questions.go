package assignment

import (
	"errors"
	"fmt"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CyclicRotate(A []int, k int) ([]int, error) { //[1] , k=0 []
	// a=[3,8,9,7,6]; k = 6
	if k < 0 {
		return nil, errors.New("K can't be negative")
	}

	if k == 0 || len(A) == 0 || len(A) == 1 {
		return A, nil
	}
	length := len(A) //length = 5

	if k > length {
		k = k % length // k = 6%5 = 6/5=1
	}
	//Aa[lower:upper] //lower included // upper not included
	//lhs = A[5-3:5] = A[2:5] = [9, 7, 6]
	lhs := A[length-k : length]
	fmt.Println("value of lhs=", lhs)
	return append(lhs, A[0:length-k]...), nil
	//A[0:5-3]=A[0;2]=[3, 8]
	//A... = 3, 8
	//APPEND([9, 7, 6], 3, 8) = [9, 7, 6, 3, 8]

}
