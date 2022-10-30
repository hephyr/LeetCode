/*
 * @lc app=leetcode id=1293 lang=typescript
 *
 * [1293] Shortest Path in a Grid with Obstacles Elimination
 */

// @lc code=start
function shortestPath(grid: number[][], k: number): number {
    const m = grid.length;
    const n = grid[0].length;
    const visited = Array(m).fill(0).map(() => Array(n).fill(0).map(() => Array(k+1).fill(false)));

    const Q = [[0, 0, k]];
    let result = 0;
    while (Q.length > 0) {
        let size = Q.length;
        while (size-- > 0) {
            const item = Q.shift();
            const x = item[0];
            const y = item[1];
            const obs = item[2];
            if (x == m-1 && y == n-1 && obs >= 0) return result;
            if (obs < 0 || visited[x][y][obs]) continue;
            visited[x][y][obs] = true;
            // UP
            if (x - 1 >= 0) {
                Q.push([x-1, y, obs - grid[x-1][y]]);
            }
            // DOWN
            if (x + 1 < m) {
                Q.push([x+1, y, obs - grid[x+1][y]]);
            }
            // RIGHT
            if (y + 1 < n) {
                Q.push([x, y+1, obs - grid[x][y+1]]);
            }
            // LEFT
            if (y - 1 >= 0) {
                Q.push([x, y-1, obs - grid[x][y-1]]);
            }
        }
        result++;
    }
    return -1;
};
// @lc code=end

