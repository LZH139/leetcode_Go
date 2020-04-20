package ArithmeticSlices

func numberOfArithmeticSlices(A []int) int {
    sum := 0
    dp :=0
    ALen := len(A)
    for i:=2;i<ALen;i++ {
        if A[i]-A[i-1] == A[i-1]-A[i-2] {
            dp +=1
            sum+=dp
        }else {
            dp = 0
        }
    }
    return sum
}