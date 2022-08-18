// https://leetcode.com/problems/merge-nodes-in-between-zeros/

// O(n) with conditions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
    original := head
    for head.Next != nil {
        curr := head.Next
        if head.Val == 0 {
            for curr.Val != 0 {
                head.Val += curr.Val
                if curr.Next.Next == nil {
                    head.Next = nil
                    return original
                }
                curr = curr.Next
            }
        }
        head.Next = curr
        head = head.Next
    }
    return original
}
