func reverseString(s []byte) {
\tletf, right := 0, len(s)-1

\tfor letf < right {
\t\ts[letf], s[right] = s[right], s[letf]
\t\tletf++
\t\tright--
\t}
}