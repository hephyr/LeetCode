/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    } else {
        h := reverseList(head.Next)
        temp := h
        for temp.Next != nil {
            temp = temp.Next
        }
        temp.Next = head
        head.Next = nil
        return h
    }
}