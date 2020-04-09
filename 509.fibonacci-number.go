/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(n int) int {
    a, b := 0, 1
    for n > 0 {
        a, b = b, a + b
        n--
    }
    return a
}
// @lc code=end

