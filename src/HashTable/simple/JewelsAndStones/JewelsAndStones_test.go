package JewelsAndStones

import "testing"

func Test_numJewelsInStones(t *testing.T) {
	type args struct {
		J string
		S string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test1",
			args:args{
				J:"aA",
				S:"aAAbbbb",
			},
			want:3,
		},

		{
			name:"test1",
			args:args{
				J:"z",
				S:"ZZ",
			},
			want:0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numJewelsInStones(tt.args.J, tt.args.S); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}