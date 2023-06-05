package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}

	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}

	current := prev.Next
	for i := 0; i < right-left; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3,
		Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	head = reverseBetween(head, 2, 4)
	for head != nil {
		print(head.Val, " ")
		head = head.Next
	}
}
