/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
    result := []int{}
    m1 := make(map[int]int)
    m2 := make(map[int]int)
    for _, v := range nums1 {
        m1[v]++
    }
    for _, v := range nums2 {
        m2[v]++
    }
    for k1, v1 := range m1 {
        if v2, ok := m2[k1]; ok {
			n := v1
            if v1 > v2 {
				n = v2
			}
            for ; n > 0; n-- {
                result = append(result, k1)
            }
        }
    }
    return result
}
// @lc code=end

