/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */

// @lc code=start
func numIslands(grid [][]byte) int {
    result := 0
    dp := make([][]bool, len(grid), len(grid))
    for i, _ := range grid {
        dp[i] = make([]bool, len(grid[i]), len(grid[i]))
    }
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == '1' && !dp[i][j] {
                result++
                dfs(i, j, grid, dp)
            }
        }
    }
    return result
}

func dfs(i, j int, grid [][]byte, dp [][]bool) {
    dp[i][j] = true
    if i - 1 >= 0 && grid[i-1][j] == '1' && !dp[i-1][j] {
        dfs(i-1, j, grid, dp)
    }
    if i + 1 < len(grid) && grid[i+1][j] == '1' && !dp[i+1][j] {
        dfs(i+1, j, grid, dp)
    }
    if j - 1 >= 0 && grid[i][j-1] == '1' && !dp[i][j-1] {
        dfs(i, j-1, grid, dp)
    }
    if j+1 < len(grid[i]) && grid[i][j+1] == '1' && !dp[i][j+1] {
        dfs(i, j+1, grid, dp)
    }
}
// @lc code=end

