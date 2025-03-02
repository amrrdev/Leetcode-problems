
func topKFrequent(nums []int, k int) []int {
\tfreq := make(map[int]int)
\tfor _, value := range nums {
\t\tfreq[value]++
\t}
\t_ = k
\tbucket := make([][]int, len(nums)+1)
\tfor num, count := range freq {
\t\tbucket[count] = append(bucket[count], num)
\t}

\tres := []int{}
\tfor i := len(bucket) - 1; i >= 0 && len(res) < k; i-- {
\t\tif bucket[i] != nil {
\t\t\tres = append(res, bucket[i]...)
\t\t}
\t}

\treturn res
}