package OccurrencesAfterBigram

func findOcurrences(text string, first string, second string) []string {

	j:=0
	flag := false
	lastWord := ""
	length := len(text)
	var res []string
	for i:=0;i<length;i++ {
		if text[i] == ' ' {
			if flag {
				res = append(res,text[j:i])
				flag = false
			}

			if lastWord == first && text[j:i] == second {
				flag = true
			}
			lastWord = text[j:i]
			j = i+1
		}
	}
	if flag {
		res = append(res,text[j:length])
	}

	return res
}
