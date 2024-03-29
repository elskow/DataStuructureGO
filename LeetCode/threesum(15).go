package main

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			var low int = i + 1
			var high int = len(nums) - 1
			var sum int = 0 - nums[i]
			for low < high {
				if nums[low]+nums[high] == sum {
					result = append(result, []int{nums[i], nums[low], nums[high]})
					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}
					low++
					high--
				} else if nums[low]+nums[high] < sum {
					low++
				} else {
					high--
				}
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	for _, v := range threeSum(nums) {
		for _, vv := range v {
			print(vv)
			print(" ")
		}
		println()
	}
}
