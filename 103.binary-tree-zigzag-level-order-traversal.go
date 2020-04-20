/*
 * @lc app=leetcode id=103 lang=golang
 *
 * [103] Binary Tree Zigzag Level Order Traversal
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
 func zigzagLevelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    if root == nil { return result }
    // current := []*TreeNode{}
    next := []*TreeNode{root}
    reverse := true
    for index := 0; len(next) != 0; index++ {
        result = append(result, []int{})
        current := next
        next = []*TreeNode{}
        reverse = !reverse
        for i := 0; i < len(current); i++ {
            if reverse {
                n := len(current)-1
                result[index] = append(result[index], current[n-i].Val)
            } else {
                result[index] = append(result[index], current[i].Val)
            }
            if current[i].Left != nil {
                next = append(next, current[i].Left)
            }
            if current[i].Right != nil {
                next = append(next, current[i].Right)
            }
        }
    }
    return result
}
// @lc code=end

