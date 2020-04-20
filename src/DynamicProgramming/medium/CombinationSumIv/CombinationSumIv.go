package CombinationSumIv

func combinationSum4(nums []int, target int) int {
    dp:= make([]int,target+1)
    dp[0] = 1
    for i:=1;i<=target;i++ {
        for _,value := range nums {
            if value<=i {
                dp[i] = dp[i]+dp[i-value]
            }
        }
    }
    return dp[target]
}