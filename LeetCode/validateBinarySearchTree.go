package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var prev *TreeNode
	return isValidBSTHelper(root, &prev)
}

func isValidBSTHelper(root *TreeNode, prev **TreeNode) bool {
	if root == nil {
		return true
	}

	if !isValidBSTHelper(root.Left, prev) {
		return false
	}

	if *prev != nil && (*prev).Val >= root.Val {
		return false
	}

	*prev = root

	return isValidBSTHelper(root.Right, prev)
}

func main() {
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	println(isValidBST(root))

	root = &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	println(isValidBST(root))
}
