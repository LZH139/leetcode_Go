package BestTimeToBuyAndSellStockWithCooldown

func maxProfit(prices []int) int {
    dp := make([][]int, len(prices)+1)

    days := len(prices)-1

    if len(prices) <2 {
        return 0
    }

    dp[0] = make([]int,3)
    //0->空仓 1—>持有 2->冷冻
    dp[0][0] = 0
    dp[0][1] = -prices[0]
    dp[0][2] = 0

    for i:=1;i<len(prices);i++ {
        dp[i] = make([]int,3)
        dp[i][0] = Max(dp[i-1][0],dp[i-1][2])
        dp[i][1] = Max(dp[i-1][0]-prices[i],dp[i-1][1])
        dp[i][2] = dp[i-1][1]+prices[i]
    }

    return Max(dp[days][0],dp[days][2])

}

func Max(x int,y int) int {
    if x>y {
        return x
    }
    return y
}