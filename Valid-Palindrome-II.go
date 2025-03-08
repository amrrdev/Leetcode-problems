func isPalindrome(s string, left, right int) bool {
\tfor left < right {
\t\tif s[left] != s[right] {
\t\t\treturn false
\t\t}
\t\tleft++
\t\tright--
\t}
\treturn true
}

func validPalindrome(s string) bool {
\tleft, right := 0, len(s)-1

\tfor left < right {
\t\tif s[left] != s[right] {
\t\t\treturn isPalindrome(s, left+1, right) || isPalindrome(s, left, right-1)
\t\t}
\t\tleft++
\t\tright--
\t}

\treturn true
}