package LongPressedName

import "testing"

func Test_isLongPressedName(t *testing.T) {
    type args struct {
        name  string
        typed string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
          name:"test1",
          args:args{
              name:  "alex",
              typed: "aaleex",
          },
          want:true,
        },
        {
         name:"test1",
         args:args{
             name:  "saeed",
             typed: "ssaaedd",
         },
         want:false,
        },
        {
         name:"test1",
         args:args{
             name:  "leelee",
             typed: "lleeelee",
         },
         want:true,
        },
        {
         name:"test1",
         args:args{
             name:  "laiden",
             typed: "laiden",
         },
         want:true,
        },
        {
         name:"test1",
         args:args{
             name:  "pyplrz",
             typed: "ppyypllr",
         },
         want:false,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := isLongPressedName(tt.args.name, tt.args.typed); got != tt.want {
                t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
            }
        })
    }
}