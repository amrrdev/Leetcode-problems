func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	diff := make([]int, n+1)

	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]
		if direction == 1 {
			diff[start]++
			if end+1 < n {
				diff[end+1]--
			}
		} else {
			diff[start]--
			if end+1 < n {
				diff[end+1]++
			}
		}
	}

	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	result := []byte(s)
	for i := 0; i < n; i++ {
		shift := diff[i]
		result[i] = byte((int(result[i]-'a')+shift%26+26)%26 + 'a')
	}

	return string(result)
}
