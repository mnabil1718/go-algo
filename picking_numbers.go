package main

import (
	"fmt"
	"math"
	"sort"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-picking-numbers/problem
func pickingNumbers(a []int32) int32 {

	var max int32 = 0

	sort.Slice(a, func(i, j int) bool {
		return int(a[i]) < int(a[j])
	})

	fmt.Println("sorted", a)

	for i := int32(0); i < int32(len(a))-1; i++ {

		temp_max := 1

		for j := i + 1; j < int32(len(a)); j++ {

			fmt.Println("comparing:", a[i], "with", a[j])

			if int(math.Abs(float64(a[i]-a[j]))) > 1 {

				fmt.Println("BREAK")

				break

			}

			temp_max++

		}

		if int32(temp_max) > max {
			max = int32(temp_max)
		}

	}

	return max

}

func main() {

	a := []int32{4, 6, 5, 3, 3, 1}

	value := pickingNumbers(a)
	fmt.Println(value)

}
