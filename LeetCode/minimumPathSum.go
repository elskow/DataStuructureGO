package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	var res [][]int
	for i := 0; i < m; i++ {
		res = append(res, make([]int, n))
	}
	res[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		res[i][0] = res[i-1][0] + grid[i][0]
	}
	for j := 1; j < n; j++ {
		res[0][j] = res[0][j-1] + grid[0][j]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = min(res[i-1][j], res[i][j-1]) + grid[i][j]
		}
	}
	return res[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	grid := [][]int{
		[]int{1, 3, 1},
		[]int{1, 5, 1},
		[]int{4, 2, 1},
	}
	println(minPathSum(grid))
}
