/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil { return nil }
    odd := head
    even := head.Next
    p := odd
    q := even
    for q != nil && q.Next != nil {
        p.Next = q.Next
        p = p.Next
        q.Next = q.Next.Next
        q = q.Next
    }
    p.Next = even
    return head
}
// @lc code=end

