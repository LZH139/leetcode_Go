package GuessNumberHigherOrLowerIi

import "math"

func getMoneyAmount(n int) int {
    mat := make([][]int,n+1)

    for i:=0;i<= n;i++ {
        mat[i] = make([]int,n+1)
    }

    for i:=2;i<= n;i++ {
        for j:=1;j<=n-i+1;j++ {
            max:= math.MaxInt32
            for k:=j;k<j+i-1;k++ {
                res:=k+Max(mat[j][k-1],mat[k+1][j+i-1])
                max = Min(res, max)
            }
            mat[j][j+i-1] = max
        }
    }
    return mat[1][n]
}

func Max(x int,y int) int {
    if x>y {
        return x
    }
    return y
}

func Min(x int, y int) int {
    if x < y {
        return x
    }
    return y
}