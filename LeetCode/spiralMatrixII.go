package main

func generateMatrix(n int) [][]int {
	var res [][]int
	for i := 0; i < n; i++ {
		res = append(res, make([]int, n))
	}
	var i, j, k, num int
	for k = 0; k < n/2; k++ {
		for j = k; j < n-k-1; j++ {
			num++
			res[k][j] = num
		}
		for i = k; i < n-k-1; i++ {
			num++
			res[i][n-k-1] = num
		}
		for j = n - k - 1; j > k; j-- {
			num++
			res[n-k-1][j] = num
		}
		for i = n - k - 1; i > k; i-- {
			num++
			res[i][k] = num
		}
	}
	if n%2 == 1 {
		num++
		res[n/2][n/2] = num
	}
	return res
}

func main() {
	n := 3
	for _, row := range generateMatrix(n) {
		for _, num := range row {
			print(num, " ")
		}
		println()
	}
	n = 1
	for _, row := range generateMatrix(n) {
		for _, num := range row {
			print(num, " ")
		}
		println()
	}
}
