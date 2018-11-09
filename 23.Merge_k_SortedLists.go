package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	result := &ListNode{}
	last := result
	for {
		min := 0
		for i := 1; i < len(lists); i++ {
			if lists[min] == nil {
				min = i
			} else if lists[i] != nil && lists[i].Val < lists[min].Val {
				min = i
			}
		}
		if lists[min] == nil {
			break
		}
		last.Next = lists[min]
		lists[min] = lists[min].Next
		last = last.Next
		last.Next = nil
	}
	return result.Next
}

func main() {
	a := &ListNode{Val: 1}
	b := &ListNode{Val: 2}
	c := &ListNode{Val: 3}
	a.Next = b
	b.Next = c
	lists := []*ListNode{a}
	result := mergeKLists(lists)
	r := result
	for r != nil {
		fmt.Println(r.Val)
		r = r.Next
	}
}
