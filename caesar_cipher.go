package main

import (
	"fmt"
	"unicode"
)

func caesarCipher(s string, k int32) string {
	// Write your code here
	var result string

	for i := 0; i < len(s); i++ {
		char := rune(s[i])

		if unicode.IsLetter(char) && 'a' <= char && char <= 'z' {
			char = rune((int(rune(s[i])-'a')+int(k))%26 + 'a')
		}

		if unicode.IsLetter(char) && 'A' <= char && char <= 'Z' {
			char = rune((int(rune(s[i])-'A')+int(k))%26 + 'A')
		}

		result += string(char)

	}

	return result
}

func main() {
	sample := "middle-Outz"
	fmt.Println(caesarCipher(sample, 2))
}
