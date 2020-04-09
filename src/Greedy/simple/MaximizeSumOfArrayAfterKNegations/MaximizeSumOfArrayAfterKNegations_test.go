package MaximizeSumOfArrayAfterKNegations

import "testing"

func Test_largestSumAfterKNegations(t *testing.T) {
    type args struct {
        A []int
        K int
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
         name:"test1",
         args:args{
             A: []int{4,2,3},
             K: 1,
         },
         want:5,
        },
        {
         name:"test1",
         args:args{
             A: []int{3,-1,0,2},
             K: 3,
         },
         want:6,
        },
        {
         name:"test1",
         args:args{
             A: []int{2,-3,-1,5,-4},
             K: 2,
         },
         want:13,
        },
        {
          name:"test1",
          args:args{
              A: []int{-2,5,0,2,-2},
              K: 3,
          },
          want:11,
        },
        {
           name:"test1",
           args:args{
               A: []int{9,-4,-7,9,6},
               K: 1,
           },
           want:27,
        },
        {
            name:"test1",
            args:args{
                A: []int{-8,3,-5,-3,-5,-2},
                K: 6,
            },
            want:22,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := largestSumAfterKNegations(tt.args.A, tt.args.K); got != tt.want {
                t.Errorf("largestSumAfterKNegations() = %v, want %v", got, tt.want)
            }
        })
    }
}