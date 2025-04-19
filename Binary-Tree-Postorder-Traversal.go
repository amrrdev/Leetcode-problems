/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    res = help(root, res)
    return res
}

func help(root *TreeNode, res []int) []int {
    if root == nil {
        return res
    }
    res = help(root.Left, res)
    res = help(root.Right, res)
    res = append(res, root.Val)
    return res
}