1func containsDuplicate(nums []int) bool {
2	set := make(map[int]struct{})
3
4	for _, v := range nums {
5		if _, ok := set[v]; ok {
6			return true
7		}
8		set[v] = struct{}{}
9
10	}
11
12	return false
13}