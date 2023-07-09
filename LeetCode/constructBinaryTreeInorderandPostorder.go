package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1)
}

func build(inorder []int, postorder []int, inStart int, inEnd int, postStart int, postEnd int) *TreeNode {
	if inStart > inEnd || postStart > postEnd {
		return nil
	}

	root := &TreeNode{Val: postorder[postEnd]}
	k := 0
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root.Val {
			k = i
			break
		}
	}

	root.Left = build(inorder, postorder, inStart, k-1, postStart, postStart+k-inStart-1)
	root.Right = build(inorder, postorder, k+1, inEnd, postStart+k-inStart, postEnd-1)

	return root
}

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := buildTree(inorder, postorder)
	for root != nil {
		println(root.Val)
		root = root.Left
	}

	inorder = []int{-1}
	postorder = []int{-1}

	root = buildTree(inorder, postorder)
	for root != nil {
		println(root.Val)
		root = root.Left
	}
}
