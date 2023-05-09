package main

import "fmt"

func maxArea(height []int) int {
	var result int
	var left int = 0
	var right int = len(height) - 1
	for left < right {
		var area int = min(height[left], height[right]) * (right - left)
		if area > result {
			result = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
	fmt.Println(maxArea([]int{4, 3, 2, 1, 4}))
}
