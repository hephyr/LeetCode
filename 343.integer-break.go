/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
    m := make([]int, n+1)
    m[1] = 1
    for i := 2; i <= n; i++ {
        m[i] = m[i-1]
        for j := i - 1; j > 1; j-- {
            temp := m[i-j]
            if m[i-j] < i-j {
                temp = i-j
            }
            if j * temp > m[i] {
                m[i] = j * temp
            }
        }
    }
    return m[n]
}
// @lc code=end

