package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var prev *TreeNode
	recoverTreeHelper(root, &prev)
}

func recoverTreeHelper(root *TreeNode, prev **TreeNode) {
	if root == nil {
		return
	}

	recoverTreeHelper(root.Left, prev)

	if *prev != nil && (*prev).Val >= root.Val {
		if *prev == nil {
			*prev = root
		} else {
			(*prev).Val, root.Val = root.Val, (*prev).Val
			*prev = nil
		}
	}

	recoverTreeHelper(root.Right, prev)
}

func main() {
	root := &TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}}
	recoverTree(root)
	for root != nil {
		println(root.Val)
		root = root.Right
	}

	root = &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}}
	recoverTree(root)
	for root != nil {
		println(root.Val)
		root = root.Right
	}
}
