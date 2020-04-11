package NumberOf1Bits

import "testing"

func Test_hammingWeight(t *testing.T) {
    type args struct {
        num uint32
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name:"test1",
            args:args{1},
            want:1,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := hammingWeight(tt.args.num); got != tt.want {
                t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
            }
        })
    }
}