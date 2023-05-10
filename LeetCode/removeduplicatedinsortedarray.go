package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var (
		i, j int
	)
	for j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
			j++
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
