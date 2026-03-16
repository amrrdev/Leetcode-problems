1func getConcatenation(nums []int) []int {
2	res := make([]int, len(nums)*2)
3	n := len(nums)
4
5	for i := range nums {
6		res[i] = nums[i]
7		res[i+n] = nums[i]
8	}
9
10	return res
11}
12