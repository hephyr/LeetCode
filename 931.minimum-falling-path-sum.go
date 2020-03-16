/*
 * @lc app=leetcode id=931 lang=golang
 *
 * [931] Minimum Falling Path Sum
 */

// @lc code=start
func minFallingPathSum(A [][]int) int {
    m := len(A)
    n := len(A[0])
    for i := 1; i < m; i++ {
        for j := 0; j < n; j++ {
            p := j-1
            if p < 0 { p = 0 }
            q := j+1
            if q == n { q = n-1 }
            A[i][j] = min(A[i-1][p], A[i-1][j], A[i-1][q]) + A[i][j]
        }
    }
    min := A[m-1][0]
    for i := 1; i < n; i++ {
        if A[m-1][i] < min {
            min = A[m-1][i]
        }
    }
    return min
}

func min(x, y, z int) int {
    if x > y {
        x, y = y, x
    }
    if x < z {
        return x
    } else {
        return z
    }
}
// @lc code=end

