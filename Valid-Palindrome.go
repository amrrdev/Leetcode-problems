func clean(s []byte) string {
\tj := 0
\tfor _, b := range s {
\t\tif 'a' <= b && b <= 'z' ||
\t\t\t'A' <= b && b <= 'Z' ||
\t\t\t'0' <= b && b <= '9' {
\t\t\ts[j] = b
\t\t\tj++
\t\t}
\t}
\treturn strings.ToLower(string(s[:j]))
}

func isPalindrome(s string) bool {
\tcleanS := clean([]byte(s))
\tif len(cleanS) == 0 {
\t\treturn true
\t}

\ti, j := 0, len(cleanS)-1
\tfor i < j {
\t\tif cleanS[i] != cleanS[j] {
\t\t\treturn false
\t\t}
\t\ti++
\t\tj--
\t}

\treturn true
}