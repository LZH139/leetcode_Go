package IsSubsequence

func isSubsequence(s string, t string) bool {
    left :=0
    tlen := len(t)
    slen := len(s)
    if slen == 0 {
        return true
    }
    if slen > tlen {
       return false
    }
    for i:=0;i<tlen;i++ {
        if s[left] == t[i] {
            left++
        }
        if left == slen {
            return true
        }
    }
    return false
}