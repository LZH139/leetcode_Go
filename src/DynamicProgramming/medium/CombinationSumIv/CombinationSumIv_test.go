package CombinationSumIv

import "testing"

func Test_combinationSum4(t *testing.T) {
    type args struct {
        nums   []int
        target int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name:"test1",
            args:args{
                nums:   []int{1, 2, 3},
                target: 4,
            },
            want:7,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := combinationSum4(tt.args.nums, tt.args.target); got != tt.want {
                t.Errorf("combinationSum4() = %v, want %v", got, tt.want)
            }
        })
    }
}