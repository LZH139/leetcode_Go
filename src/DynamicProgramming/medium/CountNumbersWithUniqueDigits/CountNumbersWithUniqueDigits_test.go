package CountNumbersWithUniqueDigits

import "testing"

func Test_countNumbersWithUniqueDigits(t *testing.T) {
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
            args:args{n:2},
            want:91,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := countNumbersWithUniqueDigits(tt.args.n); got != tt.want {
                t.Errorf("countNumbersWithUniqueDigits() = %v, want %v", got, tt.want)
            }
        })
    }
}