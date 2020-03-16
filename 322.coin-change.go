/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
func coinChange(coins []int, amount int) int {
    dp := []int{0}
    for i:= 1; i <= amount; i++ {
        dp = append(dp, amount+1)
        for _, coin := range coins {
            if i >= coin {
                dp[i] = min(dp[i-coin] + 1, dp[i])
            }
        }
    }
    if dp[amount] == amount + 1 {
        return -1
    } else {
        return dp[amount]
    }
}

func min(x, y int) int {
    if x > y {
        return y
    } else {
        return x
    }
}
// @lc code=end

