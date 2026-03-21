package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	// str := "was it a car or a cat I saw?"
	str := "tab a cat"

	val := isPalindrome(str)
	fmt.Println(val)
}

func isAlphaNumeric(str string) bool {
	var alphanumeric = regexp.MustCompile("^[a-zA-Z0-9_]*$")
	return alphanumeric.MatchString(str)
}

func isPalindrome(str string) bool {

	cleanStr := ""
	for _, char := range str {
		if isAlphaNumeric(string(char)) {
			cleanStr += strings.ToLower(string(char))
		}
	}

	fmt.Println("cleanStr -->", cleanStr)

	for i := 0; i < len(cleanStr); i++ {
		j := len(cleanStr) - 1 - i
		fmt.Printf("%s %s\n", string(cleanStr[i]), string(cleanStr[j]))
		if cleanStr[i] != cleanStr[j] {
			return false

		}
	}
	return true
}
