/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[0][0] == 1 {
		return 0
	}
    dp := [][]int{[]int{1}}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			dp[0] = append(dp[0], 0)
		} else {
			dp[0] = append(dp[0], dp[0][i-1])
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			dp = append(dp, []int{0})
		} else {
			dp = append(dp, []int{dp[i-1][0]})
		}	
		
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i] = append(dp[i], 0)
			} else {
				dp[i] = append(dp[i], dp[i][j-1] + dp[i-1][j])
			}
		}
	}
	return dp[m-1][n-1]
}
// @lc code=end

