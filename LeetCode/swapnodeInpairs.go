package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	var (
		cur, next *ListNode
	)
	cur = head
	for cur != nil && cur.Next != nil {
		next = cur.Next
		cur.Val, next.Val = next.Val, cur.Val
		cur = next.Next
	}
	return head
}

func main() {
	var head *ListNode
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head = swapPairs(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
