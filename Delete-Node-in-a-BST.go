/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }

    if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else {
        if root.Right == nil {
            return root.Left
        }
        if root.Left == nil {
            return root.Right
        }
        minNode := MinNode(root.Right)
        root.Val = minNode.Val
        root.Right = deleteNode(root.Right, minNode.Val)
    }

    return root
}

func MinNode(root *TreeNode) *TreeNode {
    for root.Left != nil {
        root = root.Left
    }
    
    return root
}