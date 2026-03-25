package main

import (
	"fmt"
	"slices"
)

func main() {
	str := "([{}])"
	// str := "(){}"
	val := isValid(str)
	fmt.Println(val)
}

func isValid(s string) bool {
	stack := []rune{}
	openings := []rune{'{', '(', '['}
	closings := []rune{'}', ')', ']'}

	for _, char := range s {
		if slices.Contains(openings, char) {
			stack = append(stack, char)
		}
		if slices.Contains(closings, char) {
			if len(stack) == 0 {
				return false
			}
			lenEnd := len(stack) - 1
			if char == '}' && stack[lenEnd] == '{' ||
				char == ')' && stack[lenEnd] == '(' ||
				char == ']' && stack[lenEnd] == '[' {
				stack = stack[:lenEnd]
				// fmt.Println("delete", stack)

				continue
			}
			return false
		}
		// fmt.Println("stack", stack)
	}

	if len(stack) == 0 {
		return true
	}
	// has unclosed items
	return false
}

// this approach doesnt work bc it assumes the brackets are only ever nested, we dont handle the case of sibling bracket
// options
// func isValid(s string) bool {
// 	runes := []rune(s)

// 	l, r := 0, len(runes)-1
// 	for l < r {
// 		switch {
// 		case runes[l] == '(' && runes[r] == ')':
// 			fallthrough
// 		case runes[l] == '{' && runes[r] == '}':
// 			fallthrough
// 		case runes[l] == '[' && runes[r] == ']':
// 			l++
// 			r--
// 			continue
// 		default:
// 			return false
// 		}
// 	}

// 	return true
// }
