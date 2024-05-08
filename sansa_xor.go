package main

import (
	"fmt"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-sansa-and-xor/problem
func sansaXor(arr []int32) int32 {
	var result int32 = 0
	if len(arr)%2 == 0 { // if length of array is even, XOR of all elements will be 0
		return result
	}

	for i := 0; i < len(arr); i += 2 { // only XOR the elements at even indices
		result ^= arr[i]
	}
	return result
}

func main() {

	arr := []int32{4, 5, 7, 5}

	fmt.Println(sansaXor(arr))
}
