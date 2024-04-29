package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	var whateverString string = "This is a line\nThis is anotherline\nA Third line, can you believe it?\nOh my god stop. It's the fourth now\nGoddammit stop!!\nSTOP DAMMIT!\nI'm done with you"

	// necessary to convert string to Reader object
	// because bufio.NewReader() accepts a Reader object
	input := strings.NewReader(whateverString)

	reader := bufio.NewReader(input)

	for {

		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		fmt.Println(string(line))

	}

}
