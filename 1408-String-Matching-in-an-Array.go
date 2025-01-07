func stringMatching(words []string) []string {
	var result []string

	for i, word := range words {
		for j, otherWord := range words {
			if i == j {
				continue
			}
			if strings.Contains(otherWord, word) {
				if !contains(result, word) {
					result = append(result, word)
				}
				break
			}
		}
	}

	return result
}

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
