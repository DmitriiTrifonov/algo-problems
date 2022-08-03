// https://leetcode.com/problems/root-equals-sum-of-children/

// Check the sum of left and right is equal to root.Val

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func checkTree(root *TreeNode) bool {
    return root.Val == (root.Left.Val + root.Right.Val)
}
