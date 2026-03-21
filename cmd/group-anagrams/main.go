package main

import "fmt"

func groupAnagrams(strs []string) [][]string {
	out := [][]string{}
	var letters [26]rune
	for i := 0; i < 26; i++ {
		letters[i] = 'a' + rune(i)
	}
	fmt.Println(letters)
	for _, str := range strs {
		println(str)
	}

	return out
}

func main() {
	arr := []string{"act", "pots", "tops", "cat", "stop", "hat"}
	val := groupAnagrams(arr)
	fmt.Println(val)
}
