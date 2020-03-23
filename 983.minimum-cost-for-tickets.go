/*
 * @lc app=leetcode id=983 lang=golang
 *
 * [983] Minimum Cost For Tickets
 */

// @lc code=start
func mincostTickets(days []int, costs []int) int {
	if len(days) == 0 {
		return 0
	}
	dp := []int{0}
	day := 0
	for i := 1; i < 366; i++ {
		if day == len(days) {
			return dp[i-1]
		}
		if i == days[day] {
			day++
			p := i-7
			if p < 0 {p=0}
			q := i-30
			if q < 0 {q=0}
			dp = append(dp, min(dp[i-1]+costs[0], dp[p]+costs[1], dp[q]+costs[2]))
		} else {
			dp = append(dp, dp[i-1])
		}
	}
	return dp[365]
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

