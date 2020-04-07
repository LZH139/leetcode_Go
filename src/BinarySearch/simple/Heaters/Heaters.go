package Heaters

import "sort"

func findRadius(houses []int, heaters []int) int {
    min := 0
    j := 0
    sort.Ints(houses)
    sort.Ints(heaters)
    for i := 0; i < len(houses); i++ {
        for ;j + 1 < len(heaters) && abs(heaters[j+1], houses[i]) <= abs(heaters[j], houses[i]); j++ {
        }
        if abs(heaters[j], houses[i]) > min {
            min = abs(heaters[j], houses[i])
        }
    }
    return min
}

func abs(x, y int) int {
    if x < y {
        return y -x
    }
    return x - y
}