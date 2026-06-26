import "strconv"

func isPalindrome(x int) bool {
    str := []rune(strconv.Itoa(x))

    n := len(str)

    for j, i := range str[0:n/2] {
        if (i != str[n-j-1]){

            return false
        }
    }

    return true
    
}