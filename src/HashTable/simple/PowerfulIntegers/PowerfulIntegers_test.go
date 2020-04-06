package PowerfulIntegers

import (
	"reflect"
	"testing"
)

func Test_powerfulIntegers(t *testing.T) {
	type args struct {
		x     int
		y     int
		bound int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name:"test1",
			args:args{
				x:     2,
				y:     3,
				bound: 10,
			},
			want:[]int{2,3,4,5,7,9,10},
		},
		{
			name:"test1",
			args:args{
				x:     3,
				y:     5,
				bound: 15,
			},
			want:[]int{2,4,6,8,10,14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := powerfulIntegers(tt.args.x, tt.args.y, tt.args.bound); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("powerfulIntegers() = %v, want %v", got, tt.want)
			}
		})
	}
}