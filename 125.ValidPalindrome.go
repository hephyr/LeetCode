func isPalindrome(s string) bool {
    if s == "" {
        return true
    }
    s = strings.ToLower(s)
    i := 0
    j := len(s) - 1
    str := []byte(s)
    for i < j {
        if !isAlphanumeric(str[i]) {
            i++
            continue
        }
        if !isAlphanumeric(str[j]) {
            j--
            continue
        }
        if str[i] != str[j] {
            return false
        }
        i++
        j--
    }
    return true
}

func isAlphanumeric(b byte) bool {
    if (b >= '0' && b <= '9') || (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
        return true
    } else {
        return false
    }
}