package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	var row, col int = len(matrix), len(matrix[0])
	var left, right int = 0, row*col - 1
	for left <= right {
		mid := (left + right) / 2
		if matrix[mid/col][mid%col] == target {
			return true
		} else if matrix[mid/col][mid%col] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34}}
	target := 3
	println(searchMatrix(matrix, target))
	target = 13
	println(searchMatrix(matrix, target))
}
