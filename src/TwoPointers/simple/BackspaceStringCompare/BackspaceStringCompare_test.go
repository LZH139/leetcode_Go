package BackspaceStringCompare

import "testing"

func Test_backspaceCompare(t *testing.T) {
    type args struct {
        S string
        T string
    }
    tests := []struct {
        name string
        args args
        want bool
    }{
        {
          name:"test1",
          args:args{
              S: "ab#c",
              T: "ad#c",
          },
          want:true,
        },
        {
          name:"test1",
          args:args{
              S: "ab##",
              T: "c#d#",
          },
          want:true,
        },
        {
          name:"test1",
          args:args{
              S: "a##c",
              T: "#a#c",
          },
          want:true,
        },
        {
          name:"test1",
          args:args{
              S: "a#c",
              T: "b",
          },
          want:false,
        },
        {
           name:"test1",
           args:args{
               S: "bxj##tw",
               T: "bxj###tw",
           },
           want:false,
        },
        {
            name:"test1",
            args:args{
                S: "a##c",
                T: "#a#c",
            },
            want:true,
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := backspaceCompare(tt.args.S, tt.args.T); got != tt.want {
                t.Errorf("backspaceCompare() = %v, want %v", got, tt.want)
            }
        })
    }
}