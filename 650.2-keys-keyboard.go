/*
 * @lc app=leetcode id=650 lang=golang
 *
 * [650] 2 Keys Keyboard
 */

// @lc code=start
func minSteps(n int) int {
    dp := []int{0, 0}
    for i := 2; i <= n; i++ {
        dp = append(dp, i)
    }
    for i := 2; i <= n; i++ {
        for j := 2; j*i <= n; j++ {
            dp[j*i] = min(dp[j*i], dp[i]+j)
        }
    }
    return dp[n]
}

func min(x, y int) int {
    if x > y {
        return y
    } else {
        return x
    }
}
// @lc code=end

