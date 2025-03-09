func removeDuplicates(nums []int) int {
\tptr := 0
\tfor i := 1; i < len(nums); i++ {
\t\tif nums[i] != nums[ptr] {
\t\t\tnums[ptr+1] = nums[i]
\t\t\tptr++
\t\t}
\t}
\treturn ptr + 1
}