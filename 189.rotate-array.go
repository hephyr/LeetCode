/*
 * @lc app=leetcode id=189 lang=golang
 *
 * [189] Rotate Array
 */

// @lc code=start
func rotate(nums []int, k int)  {
    k %= len(nums)
    reverseArray(nums, 0, len(nums) - k -1)
    reverseArray(nums, len(nums) - k, len(nums) - 1)
    reverseArray(nums, 0, len(nums) - 1)
}

func reverseArray(nums []int, begin, end int) {
    for begin < end {
        nums[begin], nums[end] = nums[end], nums[begin]
        begin++
        end--
    }
}
// @lc code=end

