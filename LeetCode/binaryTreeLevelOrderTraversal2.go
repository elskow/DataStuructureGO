package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	memo := make(map[int][]int)
	return traverseBottom(root, 0, memo)
}

func traverseBottom(root *TreeNode, level int, memo map[int][]int) [][]int {
	if root == nil {
		return nil
	}

	memo[level] = append(memo[level], root.Val)
	traverseBottom(root.Left, level+1, memo)
	traverseBottom(root.Right, level+1, memo)

	var res [][]int
	for i := len(memo) - 1; i >= 0; i-- {
		res = append(res, memo[i])
	}

	return res
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	for _, level := range levelOrderBottom(root) {
		for _, val := range level {
			print(val, " ")
		}
		println()
	}
}
