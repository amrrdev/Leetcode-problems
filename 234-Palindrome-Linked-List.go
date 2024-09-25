/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil {
        return  false
    }
    var nums []int
    current := head
    for current != nil {
        nums = append(nums,current.Val)
        current = current.Next
    }



    j:= len(nums) - 1
    for i:=0; i < len(nums) / 2; i++ {
        if nums[i] != nums[j] {
            return false
        }
        j--;
    }
    return true
}