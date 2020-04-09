package SparseArraySearchLcci

func findString(words []string, s string) int {
    i:=0
    j:=len(words)-1
    for words[i]=="" || words[j]=="" {
        if words[i] == "" {
            i++
        }
        if words[j] == "" {
            j--
        }
    }
    if helper(words[i],s) == -1 && helper(words[j],s) == 1 {
        for i<j {
            min:= i+(j-i)/2
            for words[min] == "" {
                min++
            }
            if min == j {
                min = j
            }
            temp :=helper(words[min],s)
            if temp == 0 {
                return min
            }else if temp == 1 {
                j = min
            }else {
                min++
                for words[min] == "" {
                    min++
                }
                i = min
            }
        }
        return -1
    }else{
        if helper(words[i],s) == 0 {
            return i
        }else if helper(words[j],s) == 0 {
            return j
        }else {
            return -1
        }
    }

}

func helper(word string, str string) int{
    i := 0
    j := 0
    w := uint8('a')
    s := uint8('a')
    for i<len(word) {
       w = word[i]
       s = str[j]
       if w == ' ' {
           i++
           continue
       }
       if s == ' ' {
           j++
           continue
       }
       if w != s {
           if w > s{
               return 1
           }else {
               return -1
           }
       }
       i++
       j++
    }
    if i == len(word) && j == len(str) {
       return 0
    }else {
       return -1
    }
}