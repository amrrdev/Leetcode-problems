func groupAnagrams(strs []string) [][]string {
\tfreq := make(map[[26]int][]string)
\tfor _, innerStr := range strs {
\t\tarr := [26]int{}
\t\tfor _, char := range innerStr {
\t\t\tarr[char-'a']++
\t\t}
\t\tfreq[arr] = append(freq[arr], innerStr)
\t}

\tres := [][]string{}
\tfor _, value := range freq {
\t\tres = append(res, value)
\t}
\treturn res
}
