/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int)  {
    temp := make([]int, len(nums1), len(nums1))
    copy(temp, nums1)
    i := 0
    j := 0
    k := 0
    for i < m && j < n {
        if temp[i] < nums2[j] {
            nums1[k] = temp[i]
            i++
        } else {
            nums1[k] = nums2[j]
            j++
        }
        k++
    }
    for i < m {
        nums1[k] = temp[i]
        i++
        k++
    }
    for j < n {
        nums1[k] = nums2[j]
        j++
        k++
    }
}
// @lc code=end

