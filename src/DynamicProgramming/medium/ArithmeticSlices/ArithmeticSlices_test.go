package ArithmeticSlices

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
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
            args:args{A: []int{1, 2, 3, 4}},
            want:3,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := numberOfArithmeticSlices(tt.args.A); got != tt.want {
                t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
            }
        })
    }
}