func mergeAlternately(word1 string, word2 string) string {
	mergedString := make([]byte, 0, len(word1)+len(word2))

	first, second := 0, 0
	for first < len(word1) && second < len(word2) {
		mergedString = append(mergedString, word1[first])
		mergedString = append(mergedString, word2[second])
		first++
		second++
	}
	for first < len(word1) {
		mergedString = append(mergedString, word1[first])
		first++
	}

	for second < len(word2) {
		mergedString = append(mergedString, word2[second])
		second++
	}
	return string(mergedString)
}