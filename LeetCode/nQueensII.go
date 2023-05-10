package main

func totalNQueens(n int) int {
	var result int
	totalNQueensHelper(n, 0, []int{}, &result)
	return result
}
func totalNQueensHelper(n, row int, cols []int, result *int) {
	if row == n {
		*result++
		return
	}
	for col := 0; col < n; col++ {
		if isValid(cols, row, col) {
			cols = append(cols, col)
			totalNQueensHelper(n, row+1, cols, result)
			cols = cols[:len(cols)-1]
		}
	}
}

func isValid(cols []int, row, col int) bool {
	for r, c := range cols {
		if c == col || r-c == row-col || r+c == row+col {
			return false
		}
	}
	return true
}

func main() {
	n := 4
	println(totalNQueens(n))
}
