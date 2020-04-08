package PeakIndexInAMountainArray

import "testing"

func Test_peakIndexInMountainArray(t *testing.T) {
    type args struct {
        A []int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name:"test1",
            args:args{A:[]int{0,1,0}},
            want:1,
        },
        {
            name:"test1",
            args:args{A:[]int{0,2,1,0}},
            want:1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := peakIndexInMountainArray(tt.args.A); got != tt.want {
                t.Errorf("peakIndexInMountainArray() = %v, want %v", got, tt.want)
            }
        })
    }
}