1func longestCommonPrefix(strs []string) string {
2	count := 0
3	firstPrefix := strs[0]
4	minPrefix := len(firstPrefix)
5
6	for i := 1; i < len(strs); i++ {
7		for j, c := range strs[i] {
8			// fmt.Println("c: ", c, "firstPrefix[j]: ", firstPrefix[j])
9			if j < len(firstPrefix) && c != rune(firstPrefix[j]) {
10				break
11			}
12			// if c != rune(firstPrefix[j]) {
13			// }
14			count++
15		}
16		minPrefix = min(minPrefix, count)
17		count = 0
18	}
19
20	return firstPrefix[:minPrefix]
21}
22