package main

import (
	"fmt"
)

func numTrees(n int) int {
	memo := make(map[int]int)

	return numTreesHelper(n, memo)
}

func numTreesHelper(n int, memo map[int]int) int {
	if n == 0 {
		return 1
	}

	if result, ok := memo[n]; ok {
		return result
	}

	result := 0
	for i := 1; i <= n; i++ {
		result += numTreesHelper(i-1, memo) * numTreesHelper(n-i, memo)
	}

	memo[n] = result

	return result
}

func main() {
	n := 3
	fmt.Println(numTrees(n))
	n = 1
	fmt.Println(numTrees(n))
}
