package DiYiGeZhiChuXianYiCiDeZiFuLcof

import "testing"

func Test_firstUniqChar(t *testing.T) {
    type args struct {
        s string
    }
    tests := []struct {
        name string
        args args
        want byte
    }{
        {
            name:"test",
            args:args{s:"abaccdeff"},
            want:'b',
        },
        {
            name:"test",
            args:args{s:""},
            want:' ',
        },
        {
            name:"test",
            args:args{s:"leetcode"},
            want:'l',
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := firstUniqChar(tt.args.s); got != tt.want {
                t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
            }
        })
    }
}