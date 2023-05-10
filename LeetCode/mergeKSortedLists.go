package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var (
		res, cur *ListNode
		arr      []int
	)
	for _, list := range lists {
		for list != nil {
			arr = append(arr, list.Val)
			list = list.Next
		}
	}
	sort.Ints(arr)
	for _, v := range arr {
		if res == nil {
			res = &ListNode{v, nil}
			cur = res
		} else {
			cur.Next = &ListNode{v, nil}
			cur = cur.Next
		}
	}
	return res
}

func main() {
	var lists []*ListNode
	lists = append(lists, &ListNode{1, &ListNode{4, &ListNode{5, nil}}})
	lists = append(lists, &ListNode{1, &ListNode{3, &ListNode{4, nil}}})
	lists = append(lists, &ListNode{2, &ListNode{6, nil}})
	res := mergeKLists(lists)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}
