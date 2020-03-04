/*
 * @lc app=leetcode id=108 lang=golang
 *
 * [108] Convert Sorted Array to Binary Search Tree
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
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    n := len(nums) / 2
    root := &TreeNode{}
    root.Val = nums[n]
    root.Left = sortedArrayToBST(nums[:n])
    root.Right = sortedArrayToBST(nums[n+1:])
    return root
}
// @lc code=end

