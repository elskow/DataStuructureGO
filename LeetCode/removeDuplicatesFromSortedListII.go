package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		if head.Val == head.Next.Val {
			for head != nil && head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			head = deleteDuplicates(head.Next)
		} else {
			head.Next = deleteDuplicates(head.Next)
		}
	}
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	head = deleteDuplicates(head)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
	head = &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}
	head = deleteDuplicates(head)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
