package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	result := []int{}
	inorderTraversalHelper(root, &result)
	return result
}

func inorderTraversalHelper(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}

	inorderTraversalHelper(root.Left, result)
	*result = append(*result, root.Val)
	inorderTraversalHelper(root.Right, result)
}

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}
	for _, num := range inorderTraversal(root) {
		println(num)
	}
}
