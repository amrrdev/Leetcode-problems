1import "fmt"
2
3func getConcatenation(nums []int) []int {
4    nums = append(nums, nums...)
5    return nums
6}