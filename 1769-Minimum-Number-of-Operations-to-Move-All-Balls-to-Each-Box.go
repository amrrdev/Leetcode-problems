func minOperations(boxes string) []int {
	result := make([]int, len(boxes))

	for i := 0; i < len(boxes); i++ {
		acc := 0
		for j := 0; j < len(boxes); j++ {
			if i == j {
				continue
			}
			if boxes[j] == '1' {
				acc += int(math.Abs(float64((i - j))))
			}
		}
		result[i] = acc
	}

	return result
}
