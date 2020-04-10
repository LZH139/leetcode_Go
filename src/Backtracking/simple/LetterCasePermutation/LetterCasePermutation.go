package LetterCasePermutation

import (
    "strings"
    "unicode"
)

func letterCasePermutation(S string) []string {
    list := []string{""}
    helper(S,&list)
    return list
}

func helper(S string, L *[]string){
    if S == "" {
        return
    }

    length := len(*L)
    if unicode.IsDigit(rune(S[0])) {
        for i:=0;i<length;i++ {
            (*L)[i]+=string(S[0])
        }
    }else {
        for i:=0;i<length;i++ {
            *L = append(*L,(*L)[i])
            (*L)[i]+=strings.ToLower(string(S[0]))
        }
        for i:=length;i<len(*L);i++ {
            (*L)[i]+=strings.ToUpper(string(S[0]))
        }
    }
    if len(S)>1 {
        helper(S[1:],L)
    }else if len(S) == 1 {
        helper("",L)

    }
    return
}