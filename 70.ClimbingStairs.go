func climbStairs(n int) int {
    return climbStairsIter(1, 1, n)
}

func climbStairsIter(a, b, count int) int {
    if count == 0 {
        return b
    } else {
        return climbStairsIter(a + b, a, count - 1)
    }
}