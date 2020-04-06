package PalindromePermutationLcci

func canPermutePalindrome(s string) bool {
    var list [128]int
    length := len(s)
    for i:=0;i<length;i++ {
        list[s[i]]++
    }
    count:=0
    for i:=0;i<128;i++{
        if list[i] & 1 == 1 {
            count++
        }
    }
    return count == 1 || count == 0
}