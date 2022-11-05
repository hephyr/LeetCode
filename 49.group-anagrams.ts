/*
 * @lc app=leetcode id=49 lang=typescript
 *
 * [49] Group Anagrams
 */

// @lc code=start
function groupAnagrams(strs: string[]): string[][] {
    const result = new Map<string, string[]>();
    for (const s of strs) {
        const key = s.split('').sort().join('');
        if (result.has(key)) {
            result.get(key).push(s);
        } else {
            result.set(key, [s])
        }
    }
    return Array.from(result.values());
};
// @lc code=end

