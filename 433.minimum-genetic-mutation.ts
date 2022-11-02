/*
 * @lc app=leetcode id=433 lang=typescript
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
function minMutation(start: string, end: string, bank: string[]): number {
    if (start === end) {
        return 0;
    }
    const Q = [{'gene': start, 'bank': bank, 'mutation': 0}];
    let result = 0;
    while (Q.length > 0) {
        const item = Q.shift()
        const gene = item['gene'];
        for (let i = 0; i < item['bank'].length; i++) {
            if (canMutation(gene, item['bank'][i])) {
                if (item['bank'][i] === end) { return item['mutation'] + 1}
                console.log(item['bank']);
                const b = Array.from(item['bank']);
                b.splice(i, 1);
                Q.push({'gene': item['bank'][i], 'bank': b, 'mutation': item['mutation'] + 1});
            }
        }
    }
    return -1;
};

function canMutation(geneA: string, geneB: string): boolean {
    let haveMutation = false;
    for (let i = 0; i < geneA.length; i++) {
        if (geneA[i] !== geneB[i]) {
            if (haveMutation) { return false }
            haveMutation = true;
        }
    }
    return haveMutation;
}
// @lc code=end

