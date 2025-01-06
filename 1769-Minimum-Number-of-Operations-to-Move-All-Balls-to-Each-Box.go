func minOperations(boxes string) []int {
    n := len(boxes)
    result := make([]int, n)

    leftCount, leftSum := 0, 0
    for i := 0; i < n; i++ {
        result[i] += leftSum
        if boxes[i] == '1' {
            leftCount++
        }
        leftSum += leftCount
    }

    rightCount, rightSum := 0, 0
    for i := n - 1; i >= 0; i-- {
        result[i] += rightSum
        if boxes[i] == '1' {
            rightCount++
        }
        rightSum += rightCount
    }

    return result
}