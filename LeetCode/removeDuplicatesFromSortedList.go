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
	// head = [1,1,2]
	// output: [1,2]
	head := &ListNode{1, &ListNode{1, &ListNode{2, nil}}}
	head = deleteDuplicates(head)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
	// head = [1,1,2,3,3]
	// output: [1,2,3]
	head = &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}
	head = deleteDuplicates(head)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
