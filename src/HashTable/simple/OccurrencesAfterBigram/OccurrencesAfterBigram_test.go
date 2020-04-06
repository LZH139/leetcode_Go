package OccurrencesAfterBigram

import (
	"reflect"
	"testing"
)

func Test_findOcurrences(t *testing.T) {
	type args struct {
		text   string
		first  string
		second string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name:"test1",
			args:args{
				text:   "alice is a good girl she is a good student",
				first:  "a",
				second: "good",
			},
			want:[]string{"girl","student"},
		},
		{
			name:"test1",
			args:args{
				text:   "we will we will rock you",
				first:  "we",
				second: "will",
			},
			want:[]string{"we","rock"},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findOcurrences(tt.args.text, tt.args.first, tt.args.second); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOcurrences() = %v, want %v", got, tt.want)
			}
		})
	}
}