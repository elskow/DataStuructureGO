package main

import (
	"fmt"
	"math"
)

func threeSumClosest(nums []int, target int) int {
	var result int
	var diff int = math.MaxInt64
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				var sum int = nums[i] + nums[j] + nums[k]
				var temp int = abs(sum - target)
				if temp < diff {
					diff = temp
					result = sum
				}
			}
		}
	}
	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
