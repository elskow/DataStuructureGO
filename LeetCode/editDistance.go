package main

import (
	"math"
)

func minDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	var res [][]int
	for i := 0; i < m+1; i++ {
		res = append(res, make([]int, n+1))
	}
	for i := 0; i < m+1; i++ {
		res[i][0] = i
	}
	for j := 0; j < n+1; j++ {
		res[0][j] = j
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if word1[i-1] == word2[j-1] {
				res[i][j] = res[i-1][j-1]
			} else {
				res[i][j] = int(math.Min(float64(res[i-1][j-1]), math.Min(float64(res[i-1][j]), float64(res[i][j-1])))) + 1
			}
		}
	}
	return res[m][n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	word1 := "horse"
	word2 := "ros"
	println(minDistance(word1, word2))
	word1 = "intention"
	word2 = "execution"
	println(minDistance(word1, word2))
}
