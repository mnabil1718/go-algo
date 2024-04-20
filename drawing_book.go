package main

import (
	"fmt"
	"math"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-drawing-book/problem
func pageCount(n int32, p int32) int32 {

	var front int32 = 0
	var back int32 = 0

	// front
	if p-1 > 0 {
		front = int32(math.Ceil(float64(p-1) / 2.0))
	}

	// Back
	if n-p > 1 {
		back = int32(math.Floor(float64(n-p) / 2.0))
	}

	if back < front {
		return back
	}

	return front
}

func main() {
	min := pageCount(7, 4)
	fmt.Println(min)
}
