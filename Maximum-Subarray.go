
func maxSubArray(nums []int) int {
\tcurSum, maxSum := nums[0], nums[0]
\tfor i := 1; i < len(nums); i++ {
\t\tif curSum+nums[i] > nums[i] {
\t\t\tcurSum += nums[i]
\t\t} else {
\t\t\tcurSum = nums[i]
\t\t}
\t\tmaxSum = max(maxSum, curSum)
\t}
\treturn maxSum
}