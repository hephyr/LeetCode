/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
    l, r := 0, len(nums) - 1
    for l <= r {
        mid := (l+r) / 2
        if nums[mid] == target {
            return mid
        }
        if nums[l] == nums[r] {
            return -1
        }
        if nums[l] < nums[mid] && target < nums[mid] && target >= nums[l] {
            r = mid-1
        } else if nums[r] > nums[mid] && target <= nums[r] && target > nums[mid] {
            l = mid+1
        } else if nums[l] <= nums[mid] {
            l = mid+1
        } else if nums[r] >= nums[mid] {
            r = mid-1
        }
        fmt.Println(l, r)
    }
    return -1
}
// @lc code=end

