package MinimumSubsequenceInNonIncreasingOrder

import (
    "reflect"
    "testing"
)

func Test_minSubsequence(t *testing.T) {
    type args struct {
        nums []int
    }
    tests := []struct {
        name string
        args args
        want []int
    }{
        {
            name:"test1",
            args:args{nums:[]int{4,3,10,9,8}},
            want:[]int{10,9},
        },
        {
            name:"test1",
            args:args{nums:[]int{4,4,7,6,7}},
            want:[]int{7,7,6},
        },
        {
            name:"test1",
            args:args{nums:[]int{6}},
            want:[]int{6},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := minSubsequence(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
                t.Errorf("minSubsequence() = %v, want %v", got, tt.want)
            }
        })
    }
}