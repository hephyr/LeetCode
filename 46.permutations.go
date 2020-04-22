/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
    result := [][]int{[]int{}}
    for _, v := range nums {
        for i, _ := range result {
            for j := 0; j < len(result[i]); j++ {
                temp := make([]int, len(result[i])+1)
                copy(temp, result[i][0:j])
                temp[j] = v
                copy(temp[j+1:], result[i][j:])
                result = append(result, temp)
            }
            result[i] = append(result[i], v)
        }
    }
    return result
}
// @lc code=end

