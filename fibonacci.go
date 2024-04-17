package main

import "fmt"

func fibo(param int) []int {
	sequence := []int{0, 1}
	for i := 1; i < param; i++ {
		sequence = append(sequence, sequence[len(sequence)-1]+sequence[len(sequence)-2])
	}

	return sequence
}

func main() {

	fmt.Println(fibo(6))

}
