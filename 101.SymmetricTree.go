/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    l := []*int{}
    r := []*int{}
    traversalLeft(root.Left, &l)
    traversalRight(root.Right, &r)
    if len(l) != len(r) {
        return false
    }

    for i := 0; i < len(l); i++ {
        if l[i] != nil && r[i] != nil && *r[i] != *l[i]{
            return false
        }
        if (l[i] == nil || r[i] == nil) && l[i] != r[i] {
            return false
        }
    }
    return true
}

func traversalLeft(root *TreeNode, values *[]*int) {
    if root == nil {
        *values = append(*values, nil)
        return
    } else {
        *values = append(*values, &root.Val)
    }
    traversalLeft(root.Left, values)
    traversalLeft(root.Right, values)
}

func traversalRight(root *TreeNode, values *[]*int) {
    if root == nil {
        *values = append(*values, nil)
        return
    } else {
        *values = append(*values, &root.Val)
    }
    traversalRight(root.Right, values)
    traversalRight(root.Left, values)
}