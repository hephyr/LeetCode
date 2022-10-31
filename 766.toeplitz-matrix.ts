/*
 * @lc app=leetcode id=766 lang=typescript
 *
 * [766] Toeplitz Matrix
 */

// @lc code=start
function isToeplitzMatrix(matrix: number[][]): boolean {
    const m = matrix.length;
    const n = matrix[0].length;
    const Q = [[m-1, 0, matrix[m-1][0]]];
    while (Q.length > 0) {
        let size = Q.length;
        const v = Q[0][2];
        while (size-- > 0) {
            const item = Q.shift();
            if (v !== item[2]) return false;
            const x = item[0];
            const y = item[1];
            if (y !== n-1) {
                Q.push([x, y+1, matrix[x][y+1]]);
            }
            if (y === 0 && x !== 0) {
                Q.push([x-1, 0, matrix[x-1][0]]);
            }
        }
    }
    return true;
};
// @lc code=end

