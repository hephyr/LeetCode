/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
 func rightSideView(root *TreeNode) []int {
    result := []int{}
    if root == nil {
        return result
    }
    queue := []*TreeNode{root}
    n := 1
    for len(queue) != 0 {
        n--
        if queue[0].Left != nil {
            queue = append(queue, queue[0].Left)
        }
        if queue[0].Right != nil {
            queue = append(queue, queue[0].Right)
        }
        if n == 0 {
            result = append(result, queue[0].Val)
            n = len(queue) - 1
        }
        queue = queue[1:]
    }
    return result
}
// @lc code=end

