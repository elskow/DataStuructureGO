package main

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	var res [][]int
	for i := 0; i < m; i++ {
		res = append(res, make([]int, n))
	}
	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		res[i][0] = 1
	}
	for j := 0; j < n; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		res[0][j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = res[i-1][j] + res[i][j-1]
			if obstacleGrid[i][j] == 1 {
				res[i][j] = 0
			}
		}
	}
	return res[m-1][n-1]
}

func main() {
	obstacleGrid := [][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}
	println(uniquePathsWithObstacles(obstacleGrid))
}
