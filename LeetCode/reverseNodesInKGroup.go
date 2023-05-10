package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var (
		cur, next *ListNode
		count     int
	)
	cur = head
	for cur != nil && count < k {
		cur = cur.Next
		count++
	}
	if count == k {
		cur = reverseKGroup(cur, k)
		for count > 0 {
			next = head.Next
			head.Next = cur
			cur = head
			head = next
			count--
		}
		head = cur
	}
	return head
}

func main() {
	var (
		head *ListNode
		k    int
	)
	head = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	k = 2
	head = reverseKGroup(head, k)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
