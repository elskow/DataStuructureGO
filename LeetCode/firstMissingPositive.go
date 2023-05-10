package main

import (
	"fmt"
	"sort"
)

func firstMissingPositive(nums []int) int {
	sort.Ints(nums)
	result := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == result {
			result++
		}
	}
	return result
}

func main() {
	nums := []int{1, 2, 0}
	fmt.Println(firstMissingPositive(nums))
}
