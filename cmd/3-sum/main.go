package main

import (
	"fmt"
	"slices"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	val := threesum(nums)
	fmt.Println(val)
}

func threesum(nums []int) [][]int {
	res := map[[3]int]struct{}{}
	slices.Sort(nums)

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					res[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				}
			}
		}
	}

	var result [][]int
	for triplet := range res {
		result = append(result, []int{triplet[0], triplet[1], triplet[2]})
	}

	return result
}
