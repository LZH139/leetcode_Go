package PowerOfTwo

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
    type args struct {
        n int
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
            name:"test1",
            args:args{n:1},
            want:true,
        },
        {
            name:"test1",
            args:args{n:16},
            want:true,
        },
        {
            name:"test1",
            args:args{n:218},
            want:false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := isPowerOfTwo(tt.args.n); got != tt.want {
                t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
            }
        })
    }
}