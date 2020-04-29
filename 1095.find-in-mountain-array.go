/*
 * @lc app=leetcode id=1095 lang=golang
 *
 * [1095] Find in Mountain Array
 */

// @lc code=start
/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 * type MountainArray struct {
 * }
 *
 * func (this *MountainArray) get(index int) int {}
 * func (this *MountainArray) length() int {}
 */

 func findInMountainArray(target int, mountainArr *MountainArray) int {
    low, high := 0, mountainArr.length()-1
    for low + 1 < high {
        mid := (low + high) / 2
        if mountainArr.get(mid) < mountainArr.get(mid+1) {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }
    max_index := 0
    if mountainArr.get(low) < mountainArr.get(high) {
        max_index = high
    } else {
        max_index = low
    }
    
    result := -1
    low, high = 0, max_index
    for low <= high {
        mid := (low + high) / 2
        n := mountainArr.get(mid)
        if n == target {
            return mid
        } else if n < target {
            low = mid+1
        } else {
            high = mid - 1
        }
    }
    low, high = max_index, mountainArr.length()-1
    for low <= high {
        mid := (low + high) / 2
        n := mountainArr.get(mid)
        if n == target {
            return mid
        } else if n < target {
            high = mid - 1
        } else {
            low = mid+1
        }
    }
    return result
}
// @lc code=end

