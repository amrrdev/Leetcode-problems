
func twoSum(numbers []int, target int) []int {
\tleft, right := 0, len(numbers)-1
\tfor left < right {
\t\tif numbers[left]+numbers[right] > target {
\t\t\tright--
\t\t} else if numbers[left]+numbers[right] < target {
\t\t\tleft++
\t\t} else {
\t\t\treturn []int{left + 1, right + 1}
\t\t}
\t}
\treturn []int{}
}