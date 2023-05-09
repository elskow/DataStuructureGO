package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	var carry int
	for l1 != nil || l2 != nil {
		var sum int
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum += carry
		carry = sum / 10
		node := &ListNode{Val: sum % 10}
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}
	return head
}

func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: nil}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: nil}}
	fmt.Println(addTwoNumbers(l1, l2))
}
