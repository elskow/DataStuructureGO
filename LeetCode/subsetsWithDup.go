package main

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	result := [][]int{}
	for i := 0; i <= len(nums); i++ {
		result = append(result, subsets(nums, i)...)
	}
	return result
}

func subsets(nums []int, k int) [][]int {
	if k == 0 {
		return [][]int{{}}
	}

	result := [][]int{}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for _, subset := range subsets(nums[i+1:], k-1) {
			result = append(result, append([]int{nums[i]}, subset...))
		}
	}

	return result
}

func main() {
	nums := []int{1, 2, 2}
	for _, subset := range subsetsWithDup(nums) {
		for _, num := range subset {
			print(num, " ")
		}
		println()
	}
}
