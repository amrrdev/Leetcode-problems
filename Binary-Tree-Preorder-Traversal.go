/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    res := make([]int, 0)
    res = Help(root, res)
    return res
}

func Help(root *TreeNode, res []int) []int {
    if root == nil {
        return res
    }
    res = append(res, root.Val)
    res = Help(root.Left, res)
    res = Help(root.Right, res)
    return res
}