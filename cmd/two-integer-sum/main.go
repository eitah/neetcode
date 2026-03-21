package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	target := 6

	val := twoSum(nums, target)
	fmt.Println(val)
}

func twoSum(numbers []int, target int) []int {
	// checked := map[int]bool{}

	for idx, num := range numbers {
		solution := target - num
		for idx2 := len(numbers) - 1; idx2 > idx; idx2-- {
			if numbers[idx2] > solution {
				continue // no need to check items that are too big
			}
			if numbers[idx2] == solution {
				return []int{idx + 1, idx2 + 1}
			}
		}
	}

	panic("this isnt possible")
}
