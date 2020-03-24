/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
    dp := [][]int{[]int{triangle[0][0]}}
    l := len(triangle)
    for i := 1; i < l; i++ {
        dp = append(dp, []int{})
        for j := 0; j < len(triangle[i]); j++ {
            p := j-1
            if p < 0 {p = 0}
            q := j
            if j >= len(triangle[i-1]) {q = j-1}
            dp[i] = append(dp[i], triangle[i][j] + min(dp[i-1][q], dp[i-1][p]))
        }
    }
    result := dp[l-1][0]
    for i := 1; i < len(triangle[l-1]); i++ {
        if dp[l-1][i] < result {
            result = dp[l-1][i]
        }
    }
    return result
}

func min(x, y int) int {
    if x < y {
        return x
    } else {
        return y
    }
}
// @lc code=end

