func lengthOfLongestSubstring(s string) int {
\tmyMap := make(map[byte]int)
\tstart := 0
\tres := 0
\tfor i := 0; i < len(s); i++ {
\t\tif value, ok := myMap[s[i]]; ok {
\t\t\tstart = max(value + 1, start)
\t\t}

\t\tmyMap[s[i]] = i
\t\tres = max(res, i-start+1)
\t}

\treturn res
}
