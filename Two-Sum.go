1func twoSum(nums []int, target int) []int {
2	ext := make(map[int]int, len(nums))
3
4	for i, value := range nums {
5		r := target - value
6
7		if v, ok := ext[r]; !ok {
8			ext[value] = i
9		} else {
10			return []int{i, v}
11		}
12	}
13
14	return []int{}
15}