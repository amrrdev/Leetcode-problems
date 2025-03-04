func twoSum(numbers []int, target int) []int {
\tres := make([]int, 2)
\ti, j := 0, len(numbers)-1

\tfor i < j {
\t\tif numbers[i]+numbers[j] > target {
\t\t\tj--
\t\t} else if numbers[i]+numbers[j] < target {
\t\t\ti++
\t\t} else {
\t\t\tres[0], res[1] = i+1, j+1
\t\t\tbreak
\t\t}
\t}
\treturn res
}