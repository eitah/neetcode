package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-2, 0, 0, 2, 2}
	val := threesum(nums)
	fmt.Println(val)
}

// in theory this is lower complexity
func threesum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if a > 0 {
			break
		}

		if i > 0 && a == nums[i-1] {
			continue // a duplicate of the first item has alread been checked dont keep adding it to the array
		}

		l, r := i+1, len(nums)-1

		for l < r {
			threeSum := a + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			}
			if threeSum < 0 {
				l++
			}
			if threeSum == 0 {
				result = append(result, []int{a, nums[l], nums[r]})
				l++
				r--
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			}
		}
	}

	return result
}

// this is on3 complexity because of the three loops
// func threesum(nums []int) [][]int {
// 	res := map[[3]int]struct{}{}
// 	slices.Sort(nums)

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			for k := j + 1; k < len(nums); k++ {
// 				if nums[i]+nums[j]+nums[k] == 0 {
// 					res[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
// 				}
// 			}
// 		}
// 	}

// 	var result [][]int
// 	for triplet := range res {
// 		result = append(result, []int{triplet[0], triplet[1], triplet[2]})
// 	}

// 	return result
// }
