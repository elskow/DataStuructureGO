package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	lessHead := &ListNode{}
	less := lessHead
	greaterHead := &ListNode{}
	greater := greaterHead
	for head != nil {
		if head.Val < x {
			less.Next = head
			less = less.Next
		} else {
			greater.Next = head
			greater = greater.Next
		}
		head = head.Next
	}
	greater.Next = nil
	less.Next = greaterHead.Next
	return lessHead.Next
}

func main() {
	//Input: head = [1,4,3,2,5,2], x = 3
	//	Output: [1,2,2,4,3,5]
	head := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	x := 3
	head = partition(head, x)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
