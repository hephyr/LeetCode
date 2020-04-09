/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */

// @lc code=start
func findMin(nums []int) int {
    n := len(nums) - 1
    begin, end := 0, n
    for begin < end {
        if nums[begin] < nums[end] {
            return nums[begin]
        }
        if begin + 1 == end {
            return nums[end]
        }
        mid := (begin + end) / 2
        if nums[mid] > nums[end] {
            begin = mid
        } else if nums[mid] < nums[begin] {
            end = mid
        } else if nums[begin] == nums[end] {
            end--
        }
    }
    return nums[begin]
}
// @lc code=end

