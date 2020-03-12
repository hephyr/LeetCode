/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
func uniquePaths(m int, n int) int {
	dp := [][]int{[]int{1}}
	for i := 1; i < n; i++ {
		dp[0] = append(dp[0], 1)
	}
	for i := 1; i < m; i++ {
		dp = append(dp, []int{1})
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i] = append(dp[i], dp[i][j-1] + dp[i-1][j])
		}
	}
	return dp[m-1][n-1]
}

// @lc code=end

