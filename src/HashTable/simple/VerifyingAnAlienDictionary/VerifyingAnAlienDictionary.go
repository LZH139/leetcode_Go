package VerifyingAnAlienDictionary

func isAlienSorted(words []string, order string) bool {
	mp := [26]int{}
	olength := len(order)
	for i:=0;i<olength;i++ {
		mp[order[i]-'a'] = i+1
	}

	wlength := len(words)
	j:=0
	flag := true
	lastNum := 0
	for flag {
		flag = false
		if j < len(words[0]) {
			lastNum = mp[words[0][j]-'a']
		}else{
			lastNum = 0
		}
		for i:=1;i<wlength;i++ {
			if j < len(words[i]) {
				flag = true
				if mp[words[i][j]-'a'] >= lastNum {
					lastNum = mp[words[i][j]-'a']
				}else {
					return false
				}

			}else {
				lastNum = 0
			}
			j++
		}
	}
	return true
}