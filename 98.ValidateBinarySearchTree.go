/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }

    if root.Left != nil {
        if root.Val <= root.Left.Val {
            return false
        }
        lMax := findMax(root.Left)
        if lMax >= root.Val {
            return false
        }
    }
    
    if root.Right != nil {
        if root.Val >= root.Right.Val {
            return false
        }
        rMin := findMin(root.Right)
        if rMin <= root.Val {
            return false
        }
    }
    
    if !isValidBST(root.Left) {
        return false
    }
    if !isValidBST(root.Right) {
        return false
    }
    
    return true
}

func findMax(root *TreeNode) int {
    if root.Right == nil {
        return root.Val
    } else {
        return findMax(root.Right)
    }
}

func findMin(root *TreeNode) int {
    if root.Left == nil {
        return root.Val
    } else {
        return findMin(root.Left)
    }
}