package ClimbingStairs

func waysToStep(n int) int {
    a, b, c := 1, 2, 4
    MOD := 1000000007
    for i := 0; i < n - 1; i++ {
        a, b, c = b % MOD, c % MOD, (a + b + c) % MOD
    }
    return a
}