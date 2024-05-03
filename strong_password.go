package main

import (
	"fmt"
	"strings"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-strong-password/problem
func minimumNumber(n int32, password string) int32 {
	var min int32

	if !strings.ContainsAny(password, "0123456789") {
		min++
	}

	if !strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz") {
		min++
	}

	if !strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		min++
	}

	if !strings.ContainsAny(password, "!@#$%^&*()-+") {
		min++
	}

	if int32(n)+min < 6 {
		min += 6 - int32(int32(n)+min)
	}

	return min
}

func main() {
	sample := "2bb#A"
	fmt.Println(minimumNumber(int32(len(sample)), sample))
}
