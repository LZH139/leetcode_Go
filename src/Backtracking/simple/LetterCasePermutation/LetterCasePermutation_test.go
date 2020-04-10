package LetterCasePermutation

import (
    "reflect"
    "testing"
)

func Test_letterCasePermutation(t *testing.T) {
    type args struct {
        S string
    }
    tests := []struct {
        name string
        args args
        want []string
    }{
        {
            name:"test1",
            args:args{S:"a1b2"},
            want: []string{"a1b2", "a1B2", "A1b2", "A1B2"},
        },
        {
            name:"test1",
            args:args{S:"3z4"},
            want: []string{"3z4", "3Z4"},
        },
        {
            name:"test1",
            args:args{S:"12345"},
            want: []string{"12345"},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := letterCasePermutation(tt.args.S); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
            }
        })
    }
}