package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	var n int
	for p := head; p != nil; p = p.Next {
		n++
	}
	if n == 0 {
		return head
	}
	k %= n
	if k == 0 {
		return head
	}
	var p, q *ListNode
	for p, q = head, head; k > 0; k-- {
		q = q.Next
	}
	for q.Next != nil {
		p = p.Next
		q = q.Next
	}
	res := p.Next
	p.Next = nil
	q.Next = head
	return res
}

func main() {
	head := []int{1, 2, 3, 4, 5}
	k := 2
	for _, num := range rotateRight(head, k) {
		print(num, " ")
	}
	println()
	head = []int{0, 1, 2}
	k = 4
	for _, num := range rotateRight(head, k) {
		print(num, " ")
	}
	println()
}
