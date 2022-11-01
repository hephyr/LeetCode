/*
 * @lc app=leetcode id=1706 lang=typescript
 *
 * [1706] Where Will the Ball Fall
 */

// @lc code=start
function findBall(grid: number[][]): number[] {
    const result = [];
    const m = grid.length;
    const n = grid[0].length;
    for (let i = 0; i < n; i++) {
      let current_column = i;
      for (let j = 0; j < m; j++) {
        if (grid[j][current_column] == 1) {
          if (current_column + 1 >= n || grid[j][current_column + 1] === -1) {
              current_column = -1;
              break;
          }
          current_column++;
        } else if (grid[j][current_column] == -1) {
          if (current_column - 1 < 0 || grid[j][current_column - 1] === 1) {
              current_column = -1;
              break;
          }
          current_column--;
        }
      }
      result.push(current_column);
    }
    return result;
  }
  // @lc code=end
  