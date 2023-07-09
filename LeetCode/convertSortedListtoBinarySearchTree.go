package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {

	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	// find the middle node
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	// cut the list
	prev.Next = nil

	// build the tree
	root := &TreeNode{Val: slow.Val}
	root.Left = sortedListToBST(head)
	root.Right = sortedListToBST(slow.Next)
	return root
}

func main() {
	head := &ListNode{Val: -10}
	head.Next = &ListNode{Val: -3}
	head.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next = &ListNode{Val: 9}

	sortedListToBST(head)

	for head != nil {
		println(head.Val)
		head = head.Next
	}

	head = &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}

	sortedListToBST(head)

	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
