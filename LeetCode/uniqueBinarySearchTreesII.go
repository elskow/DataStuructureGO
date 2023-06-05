package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	if n < 1 {
		return []*TreeNode{}
	}

	return generateTreesHelper(1, n)
}

func generateTreesHelper(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	result := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := generateTreesHelper(start, i-1)
		rightTrees := generateTreesHelper(i+1, end)
		for _, leftTree := range leftTrees {
			for _, rightTree := range rightTrees {
				result = append(result, &TreeNode{Val: i, Left: leftTree, Right: rightTree})
			}
		}
	}

	return result
}

func main() {
	n := 3
	for _, root := range generateTrees(n) {
		println(root.Val)
	}
}
