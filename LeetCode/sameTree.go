package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return false
}

func main() {
	p := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	q := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	println(isSameTree(p, q))

	p = &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	q = &TreeNode{1, nil, &TreeNode{2, nil, nil}}

	println(isSameTree(p, q))
}
