/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    p1 := head.Next
    p2 := head
    for p1 != p2 && p1 != nil && p1.Next != nil {
        p1 = p1.Next.Next
        p2 = p2.Next
    }
    if p1 == p2 && (p1 != nil || p1.Next != nil) {
        return true
    } else {
        return false
    }
}