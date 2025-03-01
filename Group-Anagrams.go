func sortString(s string) string {
\trunes := []byte(s)
\tresult := make([]byte, len(s))
\tcopy(result, runes)
\tsort.Slice(result, func(i, j int) bool {
\t\treturn result[i] < result[j]
\t})

\treturn string(result)
}

func groupAnagrams(strs []string) [][]string {
\tfreq := make(map[string][]string)
\tresult := [][]string{}
\tfor _, value := range strs {
\t\tsortedString := sortString(value)
\t\tarr, ok := freq[sortedString]
\t\tif !ok {
\t\t\tnewAnagram := []string{value}
\t\t\tfreq[sortedString] = newAnagram
\t\t\tcontinue
\t\t}
\t\tarr = append(arr, value)
\t\tfreq[sortedString] = arr
\t}
\tfor _, value := range freq {
\t\tresult = append(result, value)
\t}
\treturn result
}
