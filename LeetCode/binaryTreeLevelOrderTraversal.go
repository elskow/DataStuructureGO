package main

import (
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	memo := make(map[int][]int)
	depth := 0

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		memo[depth] = append(memo[depth], node.Val)
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, depth)
	var res [][]int
	for _, v := range memo {
		res = append(res, v)
	}
	//	sort the result
	sort.Slice(res, func(i, j int) bool {
		return len(res[i]) < len(res[j])
	})
	//	sort the elements in each slice
	for _, v := range res {
		sort.Ints(v)
	}
	return res
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	for _, v := range levelOrder(root) {
		for _, v2 := range v {
			print(v2, " ")
		}
		println()
	}
}
