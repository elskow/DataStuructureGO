package main

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	maxArea := 0
	for i := 0; i < len(heights); i++ {
		minHeight := heights[i]
		for j := i; j < len(heights); j++ {
			if heights[j] < minHeight {
				minHeight = heights[j]
			}
			area := minHeight * (j - i + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {
	//Input: heights = [2,1,5,6,2,3]
	//Output: 10

	heights := []int{2, 1, 5, 6, 2, 3}
	println(largestRectangleArea(heights))

	//Input: heights = [2,4]
	//Output: 4

	heights = []int{2, 4}
	println(largestRectangleArea(heights))
}
