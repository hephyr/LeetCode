/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
    result := []int{}
    foo(root, &result)
    return result[k-1]
}

func foo(node *TreeNode, a *[]int) {
    if node == nil { return }
    foo(node.Left, a)
    *a = append(*a, node.Val)
    foo(node.Right, a)
}
// @lc code=end

