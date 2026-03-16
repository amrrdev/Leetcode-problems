1func isAnagram(s string, t string) bool {
2	if len(s) != len(t) {
3		return false
4	}
5
6	cnt := [26]int{}
7	for i := range s {
8		cnt[int(s[i]-'a')]++
9		cnt[int(t[i])-'a']--
10	}
11
12	for i := range 26 {
13		if cnt[i] != 0 {
14			return false
15		}
16	}
17
18	return true
19}