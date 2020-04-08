/*
 * @lc app=leetcode id=240 lang=golang
 *
 * [240] Search a 2D Matrix II
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 {
        return false
    }
    i, j := 0, len(matrix[0]) - 1
    for i < len(matrix) && j >= 0 {
        if matrix[i][j] == target {
            return true
        } else if matrix[i][j] > target {
            j--
        } else {
            i++
        }
    }
    return false
}
// @lc code=end

