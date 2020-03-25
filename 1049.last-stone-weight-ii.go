/*
 * @lc app=leetcode id=1049 lang=golang
 *
 * [1049] Last Stone Weight II
 */

// @lc code=start
func lastStoneWeightII(stones []int) int {
    dp := [][]int{[]int{}}
    total := 0
    for _, v := range stones {
        total += v
    }
    for i := 0; i <= total/2; i++ {
        if stones[0] > i {
            dp[0] = append(dp[0], 0)
        } else {
            dp[0] = append(dp[0], stones[0])
        }
    }
    
    for i := 1; i < len(stones); i++ {
        dp = append(dp, []int{0})
        for j := 1; j <= total/2; j++ {
            if j-stones[i] < 0 {
                dp[i] = append(dp[i], dp[i-1][j])
            } else {
                dp[i] = append(dp[i], max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[i]))
            }
        }
    }
    fmt.Println(dp)
    return total - 2*dp[len(stones)-1][total/2]
}

func max(x, y int) int {
    if x > y {
        return x
    } else {
        return y
    }
}
// @lc code=end

