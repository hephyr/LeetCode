/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    root_inorder := 0
    for i, v := range inorder {
        if v == preorder[0] {
            root_inorder = i
        }
    }
    var left, right *TreeNode = nil, nil
    if root_inorder > 0 {
        left = buildTree(preorder[1:root_inorder+1], inorder[0:root_inorder])
    }
    if root_inorder+1 < len(preorder) {
        right = buildTree(preorder[root_inorder+1:], inorder[root_inorder+1:])
    }
    node := &TreeNode{
        Val: preorder[0],
        Left: left,
        Right: right,
    }
    return node
}
// @lc code=end

