/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int)  {
    m := map[int]int{
        0: 0,
        1: 0,
        2: 0,
    }
    for _, v := range nums {
        m[v]++
    }
    for i, _ := range nums {
        for j := 0; j < 3; j++ {
            if m[j] != 0 {
                m[j]--
                nums[i] = j
                break
            }
        }
    }
}
// @lc code=end

