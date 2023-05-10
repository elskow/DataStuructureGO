package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
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
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		temp = append(temp, candidates[i])
		backtrack(candidates, target-candidates[i], i+1, temp, result)
		temp = temp[:len(temp)-1]
	}
}

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
