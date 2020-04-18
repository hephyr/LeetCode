/*
 * @lc app=leetcode id=160 lang=golang
 *
 * [160] Intersection of Two Linked Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    lenA := 0
    lenB := 0
    pA := headA
    pB := headB
    for pA != nil {
        lenA++
        pA = pA.Next
    }
    for pB != nil {
        lenB++
        pB = pB.Next
    }
    p := headA
    q := headB
    n := lenA - lenB
    if lenA < lenB {
        p, q = q, p
        n = lenB - lenA
    }
    for i := n; i > 0; i-- {
        p = p.Next
    }
    for p != nil {
        if p == q {
            return p
        } else {
            p = p.Next
            q = q.Next
        }
    }
    return nil
}
// @lc code=end

