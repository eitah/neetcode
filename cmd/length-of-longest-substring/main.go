package main

import (
	"fmt"
)

func main() {
	str := "zxyzxyz"
	val := lengthOfLongestSubstring(str)
	fmt.Println(val)
}

// this is a sliding window function, where l moves each time
// the right letter is checked if there are duplicates.
func lengthOfLongestSubstring(str string) int {
	longest := 0

	charSet := make(map[byte]bool)
	l := 0

	for r := 0; r < len(str); r++ {
		for charSet[str[r]] {
			delete(charSet, str[l])
			l += 1
		}
		charSet[str[r]] = true
		lenwindow := r - l + 1
		if lenwindow > longest {
			longest = lenwindow
		}
	}

	return longest
}

// brute force solution
// func lengthOfLongestSubstring(str string) int {
// 	longest := 0

// 	for i := 0; i < len(str); i++ {
// 		lookup := make(map[byte]bool)
// 		for j := i; j < len(str); j++ {
// 			if lookup[str[j]] {
// 				break
// 			}
// 			lookup[str[j]] = true
// 		}

// 		fmt.Println(lookup)

// 		if len(lookup) > longest {
// 			longest = len(lookup)
// 		}

// 	}
// 	return longest
// }

// this totally works but I dont have slices package so fml
// func lengthOfLongestSubstring(str string) int {
// 	longest := 0

// 	var runes []rune
// 	for _, char := range str {
// 		runes = append(runes, char)
// 	}

// 	for l := 0; l < len(str); l++ {
// 		var window []rune

// 		for r := 0; r < len(str); r++ {
// 			if !slices.Contains(window, runes[r]) {
// 				window = append(window, runes[r])
// 			}
// 			// fmt.Println(window)
// 			if len(window) > longest {
// 				longest = len(window)
// 			}
// 		}
// 	}

// 	return longest
// }
