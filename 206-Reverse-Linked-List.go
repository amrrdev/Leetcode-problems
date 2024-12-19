/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    var previous, nest *ListNode = nil, nil
    current := head
    for current != nil {
        nest = current.Next
        current.Next = previous
        previous = current
        current = nest
    }
    return previous
}