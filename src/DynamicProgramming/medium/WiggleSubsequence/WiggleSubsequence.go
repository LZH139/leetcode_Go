package WiggleSubsequence

func wiggleMaxLength(nums []int) int {
    up:=1
    down:=1
    numsLen := len(nums)
    if numsLen <2 {
        return numsLen
    }
    for i:=1;i<numsLen;i++ {
        if nums[i]>nums[i-1] {
            up = down+1
        }else if nums[i] < nums[i-1] {
            down = up+1
        }
    }
    return Max(up,down)
}

func Max(x int,y int) int {
    if x>y {
        return x
    }
    return y
}