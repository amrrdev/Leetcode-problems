func prefixCount(words []string, pref string) int {
	count := 0
	prefLength := len(pref)
	for _, word := range words {
		if len(word) >= prefLength && word[:prefLength] == pref {
			count++
		}
	}
	return count
}