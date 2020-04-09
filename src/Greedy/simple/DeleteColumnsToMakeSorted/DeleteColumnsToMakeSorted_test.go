package DeleteColumnsToMakeSorted

import "testing"

func Test_minDeletionSize(t *testing.T) {
    type args struct {
        A []string
    }
    tests := []struct {
        name string
        args args
        want int
    }{
        {
           name:"test1",
           args:args{A: []string{"cba", "daf", "ghi"}},
           want:1,
        },
        {
           name:"test1",
           args:args{A: []string{"a", "b"}},
           want:0,
        },
        {
           name:"test1",
           args:args{A: []string{"zyx", "wvu", "tsr"}},
           want:3,
        },
        {
            name:"test1",
            args:args{A: []string{"rrjk","furt","guzm"}},
            want:2,
        },

    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := minDeletionSize(tt.args.A); got != tt.want {
                t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
            }
        })
    }
}