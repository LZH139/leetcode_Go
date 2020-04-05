package UncommonWordsFromTwoSentences

import (
	"reflect"
	"testing"
)

func Test_uncommonFromSentences(t *testing.T) {
	type args struct {
		A string
		B string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name:"test1",
			args:args{
				A: "this apple is sweet",
				B: "this apple is sour",
			},
			want:[]string{"sweet","sour"},
		},

		{
			name:"test2",
			args:args{
				A: "apple apple",
				B: "banana",
			},
			want: []string{"banana"},
		},

		{
			name:"test2",
			args:args{
				A: "abcd def abcd xyz",
				B: "ijk def ijk",
			},
			want: []string{"xyz"},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uncommonFromSentences(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("uncommonFromSentences() = %v, want %v", got, tt.want)
			}
		})
	}
}