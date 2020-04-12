/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            if dfs(board, i, j, word) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, i, j int, word string) bool {
    if i >= len(board) || j >= len(board[0]) || i < 0 || j < 0 || word[0] != board[i][j] {
        return false
    }
    if len(word) == 1 {
        return true
    }
    var temp byte
    temp, board[i][j] = board[i][j], '/'
    if dfs(board, i-1, j, word[1:]) || dfs(board, i+1, j, word[1:]) || dfs(board, i, j-1, word[1:]) || dfs(board, i, j+1, word[1:]) {
        return true
    } else {
        board[i][j] = temp
        return false
    }
}
// @lc code=end

