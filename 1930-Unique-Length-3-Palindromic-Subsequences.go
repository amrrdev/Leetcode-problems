func countPalindromicSubsequence(s string) int {
	firstOccurence := make(map[byte]int)
	lastOccurence := make(map[byte]int)
	result := 0

	for i := 0; i < len(s); i++ {
		if _, exists := firstOccurence[s[i]]; !exists {
			firstOccurence[s[i]] = i
		}

		lastOccurence[s[i]] = i
	}

	for char := range firstOccurence {
		first := firstOccurence[char]
		last := lastOccurence[char]

		if first == last {
			continue
		}

		uniqueBetween := make(map[byte]struct{})

		for i := first + 1; i < last; i++ {
			uniqueBetween[s[i]] = struct{}{}
		}
		result += len(uniqueBetween)
	}

	return result
}