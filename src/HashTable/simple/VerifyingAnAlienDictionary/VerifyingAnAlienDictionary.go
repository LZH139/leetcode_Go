package VerifyingAnAlienDictionary

func isAlienSorted(words []string, order string) bool {
	mp := [26]int{}
	olength := len(order)
	for i:=0;i<olength;i++ {
		mp[order[i]-'a'] = i+1
	}

	wordslength := len(words)
	for i:=1;i<wordslength;i++ {
		if !helper(&mp,words[i-1],words[i]) {
			return false
		}
	}

	return true
}

func helper(mp *[26]int, str1 string, str2 string) bool{
	cur := 0
	str1l := len(str1)
	str2l := len(str2)

	var s1Num int
	var s2Num int

	for cur<str1l || cur<str2l {
		if cur < str1l {
			s1Num = mp[str1[cur]-'a']
		}else {
			s1Num = 0
		}

		if cur < str2l {
			s2Num = mp[str2[cur]-'a']
		}else {
			s2Num = 0
		}

		if s1Num != s2Num {
			if s1Num > s2Num {
				return false
			}else {
				return true
			}
		}
		cur++
	}
	return true
}