package main

import "fmt"

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-sock-merchant/problem
func sockMerchant(n int32, ar []int32) int32 {
	var bucket = make(map[int32]int32)
	var count int32

	for i := 0; i < int(n); i++ {
		if _, present := bucket[ar[i]]; present {
			delete(bucket, ar[i])
			count++
			continue
		}

		bucket[ar[i]] = 0
	}
	return count

}

func main() {

	var arr = []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}

	var pair = sockMerchant(int32(len(arr)), arr)
	fmt.Println("pairs:", pair)

}
