package MaximumNumberOfBalloons

import "testing"

func Test_maxNumberOfBalloons(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"test1",
			args:args{text:"nlaebolko"},
			want:1,
		},

		{
			name:"test2",
			args:args{text:"loonbalxballpoon"},
			want:2,
		},

		{
			name:"test3",
			args:args{text:"leetcode"},
			want:0,
		},

		{
			name:"test3",
			args:args{text:"balon"},
			want:0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.args.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, tt.want)
			}
		})
	}
}