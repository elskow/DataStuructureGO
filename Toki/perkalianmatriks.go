package main

import (
	"fmt"
)

func main() {
	var n, m, p int
	fmt.Scan(&n, &m, &p)
	var A [101][101]int
	var B [101][101]int
	var C [101][101]int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			fmt.Scan(&B[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			C[i][j] = 0
			for k := 0; k < m; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < p; j++ {
			fmt.Print(C[i][j], " ")
		}
		fmt.Println("")
	}
}