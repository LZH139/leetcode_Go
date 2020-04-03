package ShortestCompletingWord

import "testing"

func Test_shortestCompletingWord(t *testing.T) {
	type args struct {
		licensePlate string
		words        []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name:"test1",
			args:args{
				licensePlate : "1s3 PSt",
				words :[]string{"step", "steps", "stripe", "stepple"},
			},
			want:"steps",
		},
		{
			name:"test1",
			args:args{
				licensePlate : "1s3 456",
				words :[]string{"looks", "pest", "stew", "show"},
			},
			want:"pest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestCompletingWord(tt.args.licensePlate, tt.args.words); got != tt.want {
				t.Errorf("shortestCompletingWord() = %v, want %v", got, tt.want)
			}
		})
	}
}