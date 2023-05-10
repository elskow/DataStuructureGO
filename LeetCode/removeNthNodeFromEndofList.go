package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var (
		fast, slow *ListNode
		count      int
	)
	fast = head
	slow = head
	for fast != nil {
		fast = fast.Next
		count++
		if count > n+1 {
			slow = slow.Next
		}
	}
	if count == n {
		return head.Next
	} else {
		slow.Next = slow.Next.Next
		return head
	}
}

func main() {
	var (
		head *ListNode
		n    int
	)
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n = 2
	head = removeNthFromEnd(head, n)
	for head != nil {
		println(head.Val)
		head = head.Next
	}
}
