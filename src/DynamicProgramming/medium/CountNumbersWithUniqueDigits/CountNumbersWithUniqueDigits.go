package CountNumbersWithUniqueDigits

func countNumbersWithUniqueDigits(n int) int {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return 10
    }
    dp := make([]int,n+1)
    dp[1] = 10
    dp[2] = 81
    sum:= 10+81
    for i:=3;i<=n;i++ {
        dp[i] = dp[i-1]*(10-i+1)
        sum+=dp[i]
    }
    return sum
}