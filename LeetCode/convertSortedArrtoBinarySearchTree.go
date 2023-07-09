package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, start int, end int) *TreeNode {
	if start > end {
		return nil
	}

	mid := (start + end) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = build(nums, start, mid-1)
	root.Right = build(nums, mid+1, end)

	return root
}

func main() {
	nums := []int{1, 3}
	root := sortedArrayToBST(nums)
	for root != nil {
		println(root.Val)
		root = root.Left
	}

	nums = []int{-10, -3, 0, 5, 9}
	root = sortedArrayToBST(nums)
	for root != nil {
		println(root.Val)
		root = root.Left
	}
}
