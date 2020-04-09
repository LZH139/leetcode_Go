package DeleteColumnsToMakeSorted

func minDeletionSize(A []string) int {
    count:=0
    temp := uint8('a')
    for i:=0;i<len(A[0]);i++ {
        for j:=0;j<len(A);j++ {
            if j != 0 && temp > A[j][i]{
                count++
                break
            }
            temp = A[j][i]
        }
    }
    return count
}