func minSubArrayLen(target int, nums []int) int {
\tminLen := math.MaxInt
\tsum := 0
\tstart := 0
\tfor i := 0; i < len(nums); i++ {
\t\tsum += nums[i]
\t\tfor sum >= target {
\t\t\tminLen = min(minLen, i-start+1)
\t\t\tsum -= nums[start]
\t\t\tstart++
\t\t}
\t}

\tif minLen == math.MaxInt {
\t\treturn 0
\t}
\treturn minLen
}