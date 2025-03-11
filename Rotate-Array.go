func rotate(nums []int, k int) {
    k = k % len(nums)    
    if k == len(nums) {
        return
    }
\tslices.Reverse(nums)
\tslices.Reverse(nums[:k])
\tslices.Reverse(nums[k:])
}