/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
    result := 0
    dp := make([]int, len(nums))
    for i := 0; i < len(nums); i++ {
        dp[i] = 1
        for j := 0; j < i; j++ {
            if nums[i] > nums[j] {
                dp[i] = max(dp[i], dp[j]+1)
            }
        }
        result = max(result, dp[i])
    }
    return result
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}
// @lc code=end

