package WiggleSubsequence

import "testing"

func Test_wiggleMaxLength(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name:"test1",
            args:args{nums: []int{1,7,4,9,2,5}},
            want:6,
        },
        {
            name:"test1",
            args:args{nums: []int{1,17,5,10,13,15,10,5,16,8}},
            want:7,
        },
        {
            name:"test1",
            args:args{nums: []int{1,2,3,4,5,6,7,8,9}},
            want:2,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := wiggleMaxLength(tt.args.nums); got != tt.want {
                t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
            }
        })
    }
}