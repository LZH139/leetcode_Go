package UncommonWordsFromTwoSentences

func uncommonFromSentences(A string, B string) []string {
	mp := make(map[string]int)
	i:=0

	AB := A+" "+ B
	length := len(AB)
	for j:=0;j<length;j++ {
		if AB[j] == ' ' {
			mp[AB[i:j]]++
			i = j+1
		}
		if j == length-1 {
			mp[AB[i:j+1]]++
		}

	}

	var res []string

	for k,v := range mp {
		if v == 1 {
			res = append(res,k)
		}
	}

	return res
}