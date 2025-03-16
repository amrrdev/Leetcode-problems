func characterReplacement(s string, k int) int {
\tcount := make([]int, 26)
\tmaxLength := 0
\tleft := 0
\tfor right := 0; right < len(s); right++ {
\t\tcount[s[right]-'A']++

\t\tmaxCount := 0
\t\tfor i := 0; i < 26; i++ {
\t\t\tmaxCount = max(maxCount, count[i])
\t\t}

\t\tif (right-left+1)-maxCount > k {
\t\t\tcount[s[left]-'A']--
\t\t\tleft++
\t\t}

\t\tmaxLength = max(maxLength, right-left+1)

\t}
\treturn maxLength
}