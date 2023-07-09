package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preorder []int, inorder []int, preStart int, preEnd int, inStart int, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}

	root := &TreeNode{Val: preorder[preStart]}
	k := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			k = i
			break
		}
	}

	root.Left = build(preorder, inorder, preStart+1, preStart+k-inStart, inStart, k-1)
	root.Right = build(preorder, inorder, preStart+k-inStart+1, preEnd, k+1, inEnd)

	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree(preorder, inorder)
	for root != nil {
		println(root.Val)
		root = root.Left
	}
}
