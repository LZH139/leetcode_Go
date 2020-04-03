package JewelsAndStones

func numJewelsInStones(J string, S string) int {
	list:= ['z'-'A'+1]int{}
	count:=0

	//遍历J将其中的字母装入哈希表(数组)
	for _,value := range J {
		list[value-'A']++
	}

	//遍历S统计有多少字母在哈希表里
	for _,value := range S {
		if list[value-'A'] != 0 {
			count++
		}
	}

	return count
}
