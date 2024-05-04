package main

import "fmt"

func fibo(param int) []int {
	sequence := []int{0, 1}
	for i := 1; i < param; i++ {
		sequence = append(sequence, sequence[len(sequence)-1]+sequence[len(sequence)-2])
	}

	return sequence
}

// recursive solution
	// if n == 0 {
	// 	return 0
	// }

	// if n == 1 {
	// 	return 1
	// }

	// return Fibonacci(n-1) + Fibonacci(n-2)

func main() {

	fmt.Println(fibo(6))

}
