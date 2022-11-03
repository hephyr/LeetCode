/*
 * @lc app=leetcode id=2131 lang=typescript
 *
 * [2131] Longest Palindrome by Concatenating Two Letter Words
 */

// @lc code=start
function longestPalindrome(words: string[]): number {
    const pairsWords = pairs(words);
    const singleWords = single(words);
    for (const i of singleWords) {
        if (pairsWords.has(i)) {
            singleWords.delete(i);
        }
    }
    const pairSize = pairsWords.size * 2;
    const singleSize = singleWords.size > 0 ? 2 : 0;
    return pairSize + singleSize;
};

function single(words: string[]): Set<number> {
    const result = new Set<number>();
    for (let i = 0; i < words.length; i++) {
        if (words[i][0] === words[i][1]) {
            result.add(i);
        }
    }
    return result;
}

function pairs(words: string[]): Set<number> {
    const result = new Set<number>();
    const m = new Map();
    for (let i = 0; i < words.length; i++) {
        if (m.has(words[i])) {
            result.add(m.get(words[i]).shift());
            result.add(i);
        } else {
            const key = words[i][1] + words[i][0];
            if (m.has(key)) {
                m.get(key).push(i);
            } else {
                m.set(key, [i]);
            }
        }
    }
    return result;
}
// @lc code=end

