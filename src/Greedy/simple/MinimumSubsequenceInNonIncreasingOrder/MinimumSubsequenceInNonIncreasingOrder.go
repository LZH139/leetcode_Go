package MinimumSubsequenceInNonIncreasingOrder

import "sort"

func minSubsequence(nums []int) []int {
    sort.Ints(nums)
    i:=0
    j:=len(nums)-1
    sumi := nums[i]
    sumj := nums[j]
    var res = []int{nums[j]}
    for i !=j {
        if sumi < sumj {
            i++
            sumi+=nums[i]
        }else {
            j--
            sumj+=nums[j]
            res = append(res,nums[j])
        }
    }
    return res
}
