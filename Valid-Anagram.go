1func isAnagram(s string, t string) bool {
2    m := make([]int, 200)
3    for _, v := range s {
4        intValue := int(v)
5        m[intValue] += 1
6    }
7    for _, v := range t {
8        intValue := int(v)
9        m[intValue] -= 1
10    }
11    for _, v := range m {
12        if v != 0 {
13            return false
14        }
15    }
16    return true
17}
18