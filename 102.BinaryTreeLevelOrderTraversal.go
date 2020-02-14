/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func levelOrder(root *TreeNode) [][]int {
    result := [][]int{}
    if root == nil {
        return result
    }
    nodes := []*TreeNode{root}
    node_num := 1
    l := []int{}
    for len(nodes) != 0 {
        node_num--
        if nodes[0].Left != nil {
            nodes = append(nodes, nodes[0].Left)
        }
        if nodes[0].Right != nil {
            nodes = append(nodes, nodes[0].Right)
        }
        l = append(l, nodes[0].Val)
        nodes = nodes[1:]
        if node_num == 0 {
            node_num = len(nodes)
            result = append(result, l)
            l = []int{}
        }
    }
    return result
}