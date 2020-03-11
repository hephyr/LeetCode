/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
    INT_MAX := 2147483647
    small, mid := INT_MAX, INT_MAX
    for _, v := range nums {
        if v <= small {
            small = v
        } else if v <= mid {
            mid = v
        } else {
            return true
        }
    }
    return false
}
// @lc code=end

