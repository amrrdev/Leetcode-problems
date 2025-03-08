
func merge(nums1 []int, m int, nums2 []int, n int) {
\tfirst, second := 0, 0
\tnum1Values := append([]int{}, nums1[:m]...)

\ti := 0
\tfor first < m && second < n {
\t\tif num1Values[first] < nums2[second] {
\t\t\tnums1[i] = num1Values[first]
\t\t\tfirst++
\t\t} else {
\t\t\tnums1[i] = nums2[second]
\t\t\tsecond++
\t\t}
\t\ti++
\t}
\tfor first < m {
\t\tnums1[i] = num1Values[first]
\t\tfirst++
\t\ti++
\t}
\tfor second < n {
\t\tnums1[i] = nums2[second]
\t\tsecond++
\t\ti++
\t}
}