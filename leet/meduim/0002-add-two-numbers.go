// https://leetcode.com/problems/add-two-numbers

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    dummy := new(ListNode)
    curr := dummy
    var carry int
    for l1 != nil || l2 != nil || carry != 0 {
        var x, y int
        if l1 != nil {
            x = l1.Val
        }

        if l2 != nil {
            y = l2.Val
        }

        sum := carry + x + y
        carry = sum / 10
        curr.Next = &ListNode{
            Val: sum % 10,
        }
        curr = curr.Next
        
        if l1 != nil {
            l1 = l1.Next
        }

        if l2 != nil {
            l2 = l2.Next
        }
    }
    return dummy.Next
}

// Use the carry and make a shift from start because the start of nodes is the end of number


