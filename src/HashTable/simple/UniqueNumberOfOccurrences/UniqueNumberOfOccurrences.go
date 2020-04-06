package UniqueNumberOfOccurrences

func uniqueOccurrences(arr []int) bool {
    mp := make(map[int]int)
    mp2 :=make(map[int]int)
    for _,value := range arr {
        mp[value]++
    }

    for _,value := range mp {
        if mp2[value] > 0 {
            return false
        }
        mp2[value]++
    }
    return true
}