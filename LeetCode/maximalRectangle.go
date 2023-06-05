package main

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	maxArea := 0
	heights := make([]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(heights); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}
		area := largestRectangleArea(heights)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

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
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	println(maximalRectangle(matrix))
}
