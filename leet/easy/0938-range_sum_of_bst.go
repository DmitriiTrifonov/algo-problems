// https://leetcode.com/problems/range-sum-of-bst/description/

// Going by leaf and checking interval

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
    if root == nil {
        return 0
    }
    var current int
    if root.Val >= low && root.Val <= high {
        current = root.Val
    }
    sumLeft := rangeSumBST(root.Left, low, high)
    sumRight := rangeSumBST(root.Right, low, high)
    return current + sumLeft + sumRight
}
