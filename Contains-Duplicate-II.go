func containsNearbyDuplicate(nums []int, k int) bool {
\tk++
\tmySet := make(map[int]bool)
\tptr := 0
\tfor i := 0; i < len(nums); i++ {
\t\tif i-ptr+1 > k {
\t\t\tdelete(mySet, nums[ptr])
\t\t\tptr++
\t\t}

\t\tif _, ok := mySet[nums[i]]; ok {
\t\t\treturn true
\t\t}

\t\tmySet[nums[i]] = true
\t}

\treturn false
}
