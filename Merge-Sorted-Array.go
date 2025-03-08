
func merge(nums1 []int, m int, nums2 []int, n int) {
\ti := m
\tfor _, value := range nums2 {
\t\tnums1[i] = value
\t\ti++
\t}

\tslices.Sort(nums1)

}