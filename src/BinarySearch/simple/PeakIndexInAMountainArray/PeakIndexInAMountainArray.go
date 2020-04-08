package PeakIndexInAMountainArray

func peakIndexInMountainArray(A []int) int {
    i :=0
    for {
        if A[i]<A[i+1] {
            i++
        }else {
            break
        }
    }
    return i
}