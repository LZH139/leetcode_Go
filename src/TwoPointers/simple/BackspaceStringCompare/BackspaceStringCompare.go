package BackspaceStringCompare

func backspaceCompare(S string, T string) bool {
    var list []byte
    var list1 []byte
    count:=0

    for i:=len(S)-1;i>=0;i-- {
        if S[i] == '#' {
            count++
        }else{
            if count != 0 {
                count--
                continue
            }
            list = append(list,S[i])
        }
    }
    count = 0
    for i:=len(T)-1;i>=0;i-- {
        if T[i] == '#' {
            count++
        }else{
            if count != 0 {
                count--
                continue
            }
            list1 = append(list1,T[i])
        }
    }
    lenght := len(list)
    if lenght != len(list1) {
        return false
    }
    for i:=0;i<lenght;i++ {
        if list1[i]!=list[i] {
            return false
        }
    }

    return true

}