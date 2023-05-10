package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	result := [][]int{}
	backtrack(candidates, target, 0, []int{}, &result)
	return result
}

func backtrack(candidates []int, target int, start int, temp []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, append([]int{}, temp...))
		return
	}
	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			break
		}
		temp = append(temp, candidates[i])
		backtrack(candidates, target-candidates[i], i, temp, result)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	fmt.Println(combinationSum(candidates, target))
}
