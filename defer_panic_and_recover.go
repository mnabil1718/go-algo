package main

import (
	"fmt"
	"math/rand"
)

func generateRandomNumber() int {
	return rand.Intn(100)
}

func logging() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}

func runApp() {
	defer logging()

	num := generateRandomNumber()

	if num > 50 {
		panic(fmt.Sprintf("Error! Random number %d is greater than 50", num))
	}

	fmt.Println("Random number is:", num)
}

func main() {
	runApp()
}
