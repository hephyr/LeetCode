/*
 * @lc app=leetcode id=116 lang=golang
 *
 * [116] Populating Next Right Pointers in Each Node
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
    if root != nil {
        foo(root.Left, root.Right)
    }
	return root
}

func foo(left, right *Node) {
    if left != nil {
        left.Next = right
        foo(left.Left, left.Right)
        foo(left.Right, right.Left)
        foo(right.Left, right.Right)
    }
}
// @lc code=end

