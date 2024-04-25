package main

import (
	"fmt"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-array-left-rotation/problem
func rotateLeft(d int32, arr []int32) []int32 {
	// Write your code here
	var n []int32 = make([]int32, len(arr))
	for i := 0; i < len(arr); i++ {
		if index := i - int(d); index < 0 {
			n[len(arr)+index] = arr[i]
		} else {
			n[index] = arr[i]
		}
	}
	return n
}

func main() {

	a := []int32{1, 2, 3, 4, 5}

	rotatedSlice := rotateLeft(4, a)
	fmt.Println(rotatedSlice)

}
