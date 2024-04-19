package main

import (
	"fmt"
)

func calculateFactorial(n int) int {

	if n == 1 {
		return 1
	}

	return n * calculateFactorial(n-1)

}

func main() {

	num := calculateFactorial(3)
	fmt.Println(num)
}
