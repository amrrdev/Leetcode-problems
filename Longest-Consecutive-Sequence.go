
func longestConsecutive(nums []int) int {
\tnumSet := make(map[int]struct{})
\tfor _, value := range nums {
\t\tnumSet[value] = struct{}{}
\t}

\tmaxSequenceLen := 0

\tfor num := range numSet {
\t\tif _, ok := numSet[num-1]; !ok {
\t\t\tseq := 1
\t\t\tfor {
\t\t\t\tif _, ok := numSet[num+seq]; ok {
\t\t\t\t\tseq++
\t\t\t\t\tcontinue
\t\t\t\t}
\t\t\t\tmaxSequenceLen = max(maxSequenceLen, seq)
\t\t\t\tbreak
\t\t\t}
\t\t}
\t}

\treturn maxSequenceLen
}