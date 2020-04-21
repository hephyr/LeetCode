/*
 * @lc app=leetcode id=1248 lang=golang
 *
 * [1248] Count Number of Nice Subarrays
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
    result := 0
    odd := []int{-1}
    for i, v := range nums {
        if v % 2 != 0 {
            odd = append(odd, i)   
        }
    }
    odd = append(odd, len(nums))
    for i:=1; i + k < len(odd); i++ {
        result += (odd[i] - odd[i-1]) * (odd[i+k]-odd[i+k-1])
    }
    return result
}
// @lc code=end

