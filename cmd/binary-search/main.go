package main

import "fmt"

func main() {
	// prices := []int{-1, 0, 2, 4, 6, 8}
	prices := []int{2, 5}
	// target := 4
	target := 2
	val := search(prices, target)
	fmt.Println(val)
}

func search(nums []int, target int) int {
	out := -1

	l, r := 0, len(nums)-1

	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	for l < r {
		guessIndex := l + (r-l)/2
		guessValue := nums[guessIndex]

		if guessValue == target {
			out = guessIndex
			break
		}

		if r-l == 1 {
			if nums[r] == target {
				return r
			}
			if nums[l] == target {
				return l
			}
			// no more numbers to check, lets break
			break
		}

		if target < guessValue {
			r = guessIndex
		} else {
			l = guessIndex
		}
	}

	return out
}
