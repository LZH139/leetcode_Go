package ArrangingCoins

import "testing"

func Test_arrangeCoins(t *testing.T) {
    type args struct {
        n int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
            name:"test1",
            args:args{n:5},
            want:2,
        },
        {
            name:"test1",
            args:args{n:8},
            want:3,
        },
        {
            name:"test1",
            args:args{n:2},
            want:1,
        },
        {
            name:"test1",
            args:args{n:3},
            want:2,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := arrangeCoins(tt.args.n); got != tt.want {
                t.Errorf("arrangeCoins() = %v, want %v", got, tt.want)
            }
        })
    }
}