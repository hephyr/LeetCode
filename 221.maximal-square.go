/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
    result := 0
    dp := make([][]int, 0)
    for i := 0; i < len(matrix); i++ {
        dp = append(dp, make([]int, len(matrix[i])))
        for j := 0; j < len(matrix[i]); j++ {
            if matrix[i][j] == '0' {
                dp[i][j] = 0
            } else {
                if i == 0 || j == 0 {
                    dp[i][j] = 1
                    if result < 1 {
                        result = 1
                    }
                    continue
                }
                if matrix[i][j] == matrix[i][j-1] && matrix[i][j] == matrix[i-1][j] && matrix[i][j] ==  matrix[i-1][j-1] {
                    dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
                    if dp[i][j] > result {
                        result = dp[i][j]
                    }
                } else {
                    dp[i][j] = 1
                    if result < 1 {
                        result = 1
                    }
                }
            }
        }
    }
    return result * result
}

func min(x, y, z int) int {
    if x > y {
        x, y = y, x
    }
    if x < z {
        return x
    } else {
        return z
    }
}
// @lc code=end

