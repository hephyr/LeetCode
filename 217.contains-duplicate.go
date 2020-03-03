/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
    for _, v := range nums {
        if _, ok := m[v]; ok {
            return true
        } else {
            m[v] = true
        }
    }
    return false
}
// @lc code=end

