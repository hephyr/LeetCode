func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    minPrices := prices[0]
    maxProfit := 0
    for i := 1; i < len(prices); i++ {
        if prices[i] < minPrices {
            minPrices = prices[i]
        }
        if prices[i] - minPrices > maxProfit {
            maxProfit = prices[i] - minPrices
        }
    }
    return maxProfit
}