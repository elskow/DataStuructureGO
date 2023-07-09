package main

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var nums []int
	var nodes []*TreeNode
	inorder(root, &nums, &nodes)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		nodes[i].Val = nums[i]
	}
}

func inorder(root *TreeNode, nums *[]int, nodes *[]*TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left, nums, nodes)
	*nums = append(*nums, root.Val)
	*nodes = append(*nodes, root)
	inorder(root.Right, nums, nodes)
}

func main() {
	root := &TreeNode{1, &TreeNode{3, nil, &TreeNode{2, nil, nil}}, nil}
	recoverTree(root)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Right
	}

	root = &TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{2, nil, nil}, nil}}
	recoverTree(root)
	for root != nil {
		fmt.Println(root.Val)
		root = root.Right
	}
}
