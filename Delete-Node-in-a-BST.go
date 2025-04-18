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
        if root.Left == nil {
            return root.Right
        }
        
        if root.Right == nil {
            return root.Left
        }
        minRight := MinRight(root.Right)
        root.Val = minRight.Val
        root.Right = deleteNode(root.Right, minRight.Val)
    }
    return root
}

func MinRight(root *TreeNode) *TreeNode {
    for root.Left != nil {
        root = root.Left
    }
    return root
}