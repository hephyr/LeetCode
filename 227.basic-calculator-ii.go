/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 */

// @lc code=start
func calculate(s string) int {
    return eval(s, 0, len(s))
}

func eval(s string, begin, end int) int {
    val := 0
    for ; begin < end; begin++ {
        c := s[begin]
        if c >= '0' && c <= '9' {
            val = val*10 + (int(c) - int('0'))
            
        } else if c == '+' || c == '-' {
            e := findExpEnd(s, begin+1, end)
            i := eval(s, begin+1, e)
            fmt.Println(val, i, e)
            if c == '+' {
                val = val + i
            } else {
                val = val - i
            }
            begin = e - 1
        } else if c == '*' || c == '/' {
            e := findNumEnd(s, begin+1, end)
            i := eval(s, begin+1, e)
            fmt.Println(val, i, e)
            if c == '*' {
                val = val * i
            } else {
                val = val / i
            }
            begin = e - 1
        }
    }
    fmt.Println("Val: ", val)
    return val
}

func findExpEnd(s string, begin, end int) int {
    for ; begin < end; begin++ {
        if s[begin] == '+' || s[begin] == '-' {
            fmt.Println("break")
            break
        }
    }
    return begin
}

func findNumEnd(s string, begin, end int) int {
    flag := false
    for ; begin < end; begin++ {
        c := s[begin]
        if c >= '0' && c <= '9' {
            flag = true
        } else {
            if flag { break }
        }
    }
    return begin
}
// @lc code=end

