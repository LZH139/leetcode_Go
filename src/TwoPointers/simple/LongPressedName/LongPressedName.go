package LongPressedName

func isLongPressedName(name string, typed string) bool {
    nlen := len(name)
    tlen := len(typed)
    i := 1
    j := 1
    counti :=0
    countj := 0
    for i < nlen || j <tlen {
        for i < nlen && name[i] == name[i-1] {
            i++
            counti++
        }
        for j<tlen && typed[j] == typed[j-1] {
            j++
            countj++
        }
        if typed[j-1] != name[i-1] || countj < counti {
            return false
        }
        i++
        j++
        counti,countj = 0,0
    }
    return name[nlen-1] == typed[tlen-1]
}