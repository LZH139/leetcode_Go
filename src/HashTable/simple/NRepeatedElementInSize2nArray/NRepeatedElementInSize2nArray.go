package NRepeatedElementInSize2nArray

func repeatedNTimes(A []int) int {
	var mp [10001]int
	for _,value := range A {
		mp[value]++
		if mp[value] > 1 {
			return value
		}
	}
	return -1
}