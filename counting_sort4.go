package main

import (
	"fmt"
	"strconv"
	"strings"
)

// source: https://www.hackerrank.com/challenges/three-month-preparation-kit-countingsort4/problem
func countSort(arr [][]string) {

	var result string
	max := 0
	for i := 0; i < len(arr); i++ {
		num, _ := strconv.Atoi(arr[i][0])
		if num > max {
			max = num
		}
	}

	var container [][]string = make([][]string, max+1)

	for i := 0; i < len(arr); i++ {
		index, _ := strconv.Atoi(arr[i][0])
		if i < len(arr)/2 {
			container[index] = append(container[index], "-")
		} else {
			container[index] = append(container[index], arr[i][1])
		}
	}

	for i := 0; i < len(container); i++ {
		joined := strings.Join(container[i], " ")
		result += joined + " "

	}

	fmt.Println(strings.TrimSpace(result))
}

func main() {

	arr := [][]string{
		{"0", "ab"},
		{"6", "cd"},
		{"0", "ef"},
		{"6", "gh"},
		{"4", "ij"},
		{"0", "ab"},
		{"6", "cd"},
		{"0", "ef"},
		{"6", "gh"},
		{"0", "ij"},
		{"4", "that"},
		{"3", "be"},
		{"0", "to"},
		{"1", "be"},
		{"5", "question"},
		{"1", "or"},
		{"2", "not"},
		{"4", "is"},
		{"2", "to"},
		{"4", "the"},
	}

	countSort(arr)
}
