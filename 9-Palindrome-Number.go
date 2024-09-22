import "strconv"
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    str := strconv.Itoa(x)
    j := len(str) - 1
    for i:= 0; i< len(str) / 2; i++ {
        if str[i] != str[j] {
            return false
        }
        j--
    }
    return true
}