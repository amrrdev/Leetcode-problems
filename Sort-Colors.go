func sortColors(nums []int) {
\tl, m, r := 0, 0, len(nums)-1
\tfor m <= r {
\t\tif nums[m] == 2 {
\t\t\tnums[m], nums[r] = nums[r], nums[m]
\t\t\tr--
\t\t} else if nums[m] == 1 {
\t\t\tm++
\t\t} else {
\t\t\tnums[l], nums[m] = nums[m], nums[l]
\t\t\tl++
\t\t\tm++
\t\t}
\t}
\t// fmt.Println(nums)
}