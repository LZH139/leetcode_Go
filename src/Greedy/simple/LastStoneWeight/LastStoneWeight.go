package LastStoneWeight

import "sort"

func lastStoneWeight(stones []int) int {
    length := len(stones)
    if length == 1 {
        return stones[0]
    }
    sort.Ints(stones)
    for stones[length-2] != 0 {
        stones[length-1] = stones[length-1] - stones[length-2]
        stones[length-2] = 0
        sort.Ints(stones)
    }
    return stones[length-1]
}