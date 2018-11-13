package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	p := head
	newhead := &ListNode{Next: p}
	pre := newhead
	i := 1
	for q := head; q != nil; q = q.Next {
		if i == k {
			i = 0
			pre.Next = q
			tail := reverse(p, q)
			pre = tail
			p = pre.Next
			q = tail
		}
		i++
	}
	return newhead.Next
}

func reverse(p, q *ListNode) *ListNode {
	if p == nil {
		return nil
	}
	if p == q {
		return q
	}
	tail := reverse(p.Next, q)
	p.Next = tail.Next
	tail.Next = p
	return p
}

func main() {
	e := &ListNode{Val: 5}
	d := &ListNode{Val: 4, Next: e}
	c := &ListNode{Val: 3, Next: d}
	b := &ListNode{Val: 2, Next: c}
	a := &ListNode{Val: 1, Next: b}
	h := reverseKGroup(a, 2)
	for p := h; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}
