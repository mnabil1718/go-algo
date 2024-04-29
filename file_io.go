package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)

	if err != nil {
		return err
	}

	// defer should be used right after error checking is done
	defer file.Close()

	_, err = file.WriteString(message)
	if err != nil {
		return err
	}

	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0777)
	if err != nil {
		return "", err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var content string

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		content += string(line) + "\n"
	}

	return content, nil
}

func main() {
	var whateverString string = "This is a line\nThis is anotherline\nA Third line, can you believe it?\nOh my god stop. It's the fourth now\nGoddammit stop!!\nSTOP DAMMIT!\nI'm done with you"

	err := createNewFile("test.txt", whateverString)
	if err != nil {
		fmt.Println(err.Error())
		panic("Error creating file")
	}

	content, err := readFile("test.txt")
	if err != nil {
		fmt.Println(err.Error())
		panic("Error reading file")
	}

	fmt.Println(content)

}
