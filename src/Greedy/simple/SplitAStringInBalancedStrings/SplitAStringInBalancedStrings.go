package SplitAStringInBalancedStrings

func balancedStringSplit(s string) int {
    sum:=0
    count:=0
    length := len(s)
    for i:=0;i<length;i++ {
        if s[i] == 'R'{
            count++
        }else{
            count--
        }
        if count == 0 {
            sum++
        }
    }
    return sum
}