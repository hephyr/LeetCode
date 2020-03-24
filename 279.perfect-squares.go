/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start
func numSquares(n int) int {
    squares := []int{0, 1}
    for i := 2; i*i <= n; i++ {
        squares = append(squares, i*i)
    }
    dp := []int{0, 1}
    for i := 2; i <= n; i++ {
        dp = append(dp, i)
    }
    for i := 2; i <= n; i++ {
        for j := 1; j < len(squares); j++ {
            if i-squares[j] >= 0 {
                dp[i] = min(dp[i], dp[i-squares[j]]+1)
            }
        }
    }
    return dp[n]
}

func min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}
// @lc code=end

