/*
 * @lc app=leetcode id=234 lang=golang
 *
 * [234] Palindrome Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    p := head
    length := 0
    for p != nil {
        p = p.Next
        length++
    }
    p = head
    for i := 0; i < length/2; i++ {
        p = p.Next
    }
    if length%1 != 0 {
        p = p.Next
    }
    q := reverseList(p)
    p = head
    for i := 0; i < length/2; i++ {
        if p.Val != q.Val {
            return false
        }
        p = p.Next
        q = q.Next
    }
    return true
}

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    p := head.Next
    for p != nil {
        temp := p
        p, p.Next = p.Next, head
        head = temp
    }
    return head
}
// @lc code=end

