package main

import (
	"math"
)

func solveNQueens(n int) [][]string {
	var result [][]string
	solveNQueensHelper(n, 0, []int{}, &result)
	return result
}

func solveNQueensHelper(n, row int, col []int, result *[][]string) {
	if row == n {
		*result = append(*result, convert(col))
		return
	}
	for i := 0; i < n; i++ {
		col = append(col, i)
		if isValid(col) {
			solveNQueensHelper(n, row+1, col, result)
		}
		col = col[:len(col)-1]
	}
}

func isValid(col []int) bool {
	row := len(col) - 1
	for i := 0; i < row; i++ {
		diff := int(math.Abs(float64(col[i] - col[row])))
		if diff == 0 || diff == row-i {
			return false
		}
	}
	return true
}

func convert(col []int) []string {
	var result []string
	for _, v := range col {
		var s string
		for i := 0; i < len(col); i++ {
			if i == v {
				s += "Q"
			} else {
				s += "."
			}
		}
		result = append(result, s)
	}
	return result
}

func main() {
	n := 4
	for _, v := range solveNQueens(n) {
		for _, vv := range v {
			print(vv, " ")
		}
		println()
	}
}
