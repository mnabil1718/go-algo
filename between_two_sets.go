package main

import "fmt"

// source: https://www.hackerrank.com/challenges/between-two-sets/problem
func getTotalX(a []int32, b []int32) int32 {
	var total int32

	for i := int32(1); i <= 100; i++ {
		fmt.Println("number:", i)
		var invalid bool

		for _, element1 := range a {
			if i%element1 != 0 {
				fmt.Println("element1:", element1, "is not a factor of", i)
				invalid = true
				break
			}
		}

		if invalid {
			continue
		}

		for _, element2 := range b {
			if element2%i != 0 {
				fmt.Println("number:", i, "is not a factor of element2:", element2)
				invalid = true
				break
			}
		}

		if invalid {
			continue
		}

		total++
		fmt.Println("NUMBER:", i, "IS VALID")

	}

	return total
}

func main() {

	a := []int32{2, 4}
	b := []int32{16, 32, 96}

	value := getTotalX(a, b)
	fmt.Println(value)

}
