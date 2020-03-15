/*
 * @lc app=leetcode id=746 lang=golang
 *
 * [746] Min Cost Climbing Stairs
 */

// @lc code=start
func minCostClimbingStairs(cost []int) int {
    n := len(cost)
    for i := 2; i < n; i++ {
        cost[i] = min(cost[i-1], cost[i-2]) + cost[i]
    }
    return min(cost[n-1], cost[n-2])
}

func min(x, y int) int {
    if x > y {
        return y
    } else {
        return x
    }
}
// @lc code=end

