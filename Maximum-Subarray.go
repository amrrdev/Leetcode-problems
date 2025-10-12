func maxSubArray(arr []int) int {
	currSum := arr[0]
	maxSum := arr[0]

	for i := 1; i < len(arr); i++ {
		currSum = max(arr[i], currSum+arr[i])
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}
