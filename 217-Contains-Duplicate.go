func containsDuplicate(nums []int) bool {
\thashSet := make(map[int]struct{})
\tfor _, value := range nums {
\t\tif _, ok := hashSet[value]; ok {
\t\t\treturn true
\t\t}
\t\thashSet[value] = struct{}{}
\t}
\treturn false
}
