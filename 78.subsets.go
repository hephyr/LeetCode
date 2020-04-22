/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
func subsets(nums []int) [][]int {
    result := [][]int{[]int{}}
    for _, v := range nums {
        for _, r := range result {
            temp := make([]int, len(r)+1)
            copy(temp, r)
            temp[len(r)] = v
            result = append(result, temp)
        }
    }
    return result
}
// @lc code=end

