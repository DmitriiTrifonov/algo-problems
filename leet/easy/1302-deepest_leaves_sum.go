// https://leetcode.com/problems/deepest-leaves-sum/

// From left to right and storing it in a map

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    sums := map[int]int{}
    checkNode(root, sums, 0)
    
    max := 0
    
    for k := range sums {
        if max < k {
            max = k
        }
    }
    return sums[max]
}

func checkNode(root *TreeNode, m map[int]int, level int) {
    if root == nil {
        return
    }
    m[level] += root.Val
    checkNode(root.Left, m, level + 1)
    checkNode(root.Right, m, level + 1)
    
}
