func removeNonAlphanumeric(s string) string {
\tnewString := []byte{}
\tfor i := 0; i < len(s); i++ {
\t\tif 'a' <= s[i] && s[i] <= 'z' ||
\t\t\t'A' <= s[i] && s[i] <= 'Z' ||
\t\t\t'0' <= s[i] && s[i] <= '9' {
\t\t\tnewString = append(newString, s[i])
\t\t}
\t}

\treturn strings.ToLower(string(newString))
}

func isPalindrome(s string) bool {
\tnewString := removeNonAlphanumeric(s)
\tif len(newString) == 0 {
\t\treturn true
\t}
\tfmt.Println(len(newString))
\tleft, right := 0, len(newString)-1
\tfor left < right {
\t\tif newString[left] != newString[right] {
\t\t\treturn false
\t\t}
\t\tleft++
\t\tright--
\t}
\treturn true
}
