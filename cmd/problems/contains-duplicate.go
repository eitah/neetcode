package main

func main() {
	arr := []int{1, 2, 3, 3}
	err, val := containsDuplicate(arr)
	if err != nil {
		panic(err)
	}
	println(val)
}

func containsDuplicate(nums []int) (error, bool) {
	seen := map[int]bool{}
	for _, val := range nums {
		if seen[val] == true {
			return nil, true
		}
		seen[val] = true
	}

	return nil, false
}
