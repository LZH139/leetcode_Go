package ShortestCompletingWord

import "unicode"

func shortestCompletingWord(licensePlate string, words []string) string {
	lengthWords := len(words)
	min:= 16
	minWord := ""

	//依次遍历取出words里的每个单词
	for i:=0;i<lengthWords;i++ {
		wordLength := len(words[i])
		list := make([]int,26)

		//如果该单词比当前最短完整词的长度长则舍去
		if wordLength<min {

			//依次将每个单词转换到长度为26的哈希表里(切片做哈希表)
			for j:=0;j<wordLength;j++ {
				list[words[i][j]-'a']++
			}

			//依次遍历比较牌照中的字母是否都在该单词里
			if  compare(list,licensePlate) {

				//如果存在则更新最短完整词和最短长度
				min = wordLength
				minWord = words[i]
			}
		}
	}
	return minWord
}

func compare(list []int, licensePlate string) bool {
	for i:=0;i<len(licensePlate);i++ {
		if letter := licensePlate[i];unicode.IsLetter(rune(letter)){
			if list[(letter | 32)-'a'] <= 0 {
				return false
			}
			list[(letter | 32)-'a']--
		}
	}
	return true
}

//func shortestCompletingWord(licensePlate string, words []string) string {
//	res := ""
//	listLic := make([]int, 26)
//	n := 0
//	lengthLic := len(licensePlate)
//	lengthWord := len(words)
//	for i:=0;i<lengthLic;i++ {
//		if val:=licensePlate[i];unicode.IsLetter(rune(val)) {
//			n++
//			listLic[(val | 32) - 'a']++
//		}
//	}
//	for j:=0;j<lengthWord;j++ {
//		if len(words[j]) < n {
//			continue
//		}
//		temp := make([]int, 26)
//		for _, val := range words[j] {
//			temp[(val | 32) - 'a']++
//		}
//		i := 0
//		for ; i < 26; i++ {
//			if listLic[i] > temp[i] {
//				break
//			}
//		}
//		if i == 26 {
//			temp := len(res)
//			if temp == 0 || len(words[j]) < temp {
//				res = words[j]
//			}
//		}
//	}
//	return res
//}