package SubdomainVisitCount

import (
	"reflect"
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	type args struct {
		cpdomains []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name:"test1",
			args:args{
				//cpdomains : []string{"9001 discuss.leetcode.com"},
				cpdomains : []string{"9001 com"},
			},
			//want:[]string{"9001 leetcode.com", "9001 discuss.leetcode.com",  "9001 com"},
			want:[]string{"9001 com"},
		},
		//{
		//	name:"test1",
		//	args:args{
		//		cpdomains : []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"},
		//	},
		//	want:[]string{"5 wiki.org","951 com", "901 mail.com","900 google.mail.com", "50 yahoo.com","1 intel.mail.com","5 org"},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subdomainVisits(tt.args.cpdomains); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}