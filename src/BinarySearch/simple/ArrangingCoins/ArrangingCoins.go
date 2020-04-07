package ArrangingCoins

func arrangeCoins(n int) int {
    i:=1
    for ;i<=n;i++ {
        n -=i
    }
    return i-1
}