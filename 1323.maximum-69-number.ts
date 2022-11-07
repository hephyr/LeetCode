/*
 * @lc app=leetcode id=1323 lang=typescript
 *
 * [1323] Maximum 69 Number
 */

// @lc code=start
function maximum69Number (num: number): number {
    const s = num.toString();
    for (let i = 0; i < s.length; i++) {
        if (s[i] === '6') {
            return num + 3 * (10 ** (s.length - 1 - i));
        }
    }
    return num;
};
// @lc code=end

