package HouseRobber

func rob(nums []int) int {
    n2 := 0
    n1 := 0
    sum := 0
    for _,value := range nums {
        sum = Max(n2+value,n1)
        n2 = n1
        n1 = sum
    }
    return sum
}

func Max(x int, y int) int {
    if x>y {
        return x
    }
    return y
}