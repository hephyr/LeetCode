/*
 * @lc app=leetcode id=122 lang=golang
 *
 * [122] Best Time to Buy and Sell Stock II
 */

// @lc code=start
func maxProfit(prices []int) int {
    result := 0
    for i, j := 0, 1; j < len(prices); i, j = i+1, j+1 {
        if prices[i] < prices[j] {
            result += prices[j] - prices[i]
        }
    }
    return result
}
// @lc code=end

