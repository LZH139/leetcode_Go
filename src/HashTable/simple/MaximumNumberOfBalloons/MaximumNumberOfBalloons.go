package MaximumNumberOfBalloons

func maxNumberOfBalloons(text string) int {
	var list [26]int
	str := "balloon"

	for i:=0;i<len(text);i++ {
		list[text[i]-'a']++
	}

	res := 0
	for true {
		for i:=0;i<7;i++ {
			list[str[i]-'a']--
			if list[str[i]-'a'] < 0 {
				return res
			}
		}
		res++
	}
	return -1
}

func Min(x int ,y int) int {
	if x>y {
		return y
	}
	return x
}