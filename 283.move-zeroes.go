/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int)  {
    for i, j := 0, 0; j < len(nums); j++ {
        if nums[j] != 0 {
            nums[i], nums[j] = nums[j], nums[i]
            i++
        }
    }
}
// @lc code=end

