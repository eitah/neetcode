package main

import "fmt"

func main() {
	prices := []int{10, 8, 7, 5, 2}
	val := bestTime(prices)
	fmt.Println(val)
}

func bestTime(arr []int) int {
	profit := 0

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j]-arr[i] > profit {
				profit = arr[j] - arr[i]
			}
		}
	}

	return profit
}
