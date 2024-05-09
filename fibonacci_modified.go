package main

import (
	"fmt"
	"math/big"
)

// source: https://www.hackerrank.com/challenges/fibonacci-modified/problem
func fibonacciModified(t1, t2 int32, n int32) *big.Int {
	temp_arr := make([]*big.Int, n)
	temp_arr[0] = big.NewInt(int64(t1))
	temp_arr[1] = big.NewInt(int64(t2))

	for i := 2; i < int(n); i++ {
		temp_arr[i] = new(big.Int).Add(temp_arr[i-2], new(big.Int).Mul(temp_arr[i-1], temp_arr[i-1]))
	}

	return temp_arr[n-1]
}

func main() {

	fmt.Println(fibonacciModified(0, 1, 10))
}
