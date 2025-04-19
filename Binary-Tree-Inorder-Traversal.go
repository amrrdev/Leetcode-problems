/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
\tres := make([]int, 0)
\tres = Help(root, res)
\treturn res
}

func Help(root *TreeNode, res []int) []int {
\tif root == nil {
\t\treturn res
\t}
\tres = Help(root.Left, res)
\tres = append(res, root.Val)
\tres = Help(root.Right, res)
\treturn res
}
