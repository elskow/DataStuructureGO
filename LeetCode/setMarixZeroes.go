package main

func setZeroes(matrix [][]int) {
	var row, col []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				row = append(row, i)
				col = append(col, j)
			}
		}
	}
	for _, v := range row {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[v][i] = 0
		}
	}
	for _, v := range col {
		for i := 0; i < len(matrix); i++ {
			matrix[i][v] = 0
		}
	}
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			print(matrix[i][j], " ")
		}
		println()
	}
}
