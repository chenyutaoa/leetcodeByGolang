package leetcode

import (
	"reflect"
	"testing"
)

func Test_findRepeatedDnaSequences(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"",args{s: "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"},[]string{"AAAAACCCCC","CCCCCAAAAA"}},
		{"",args{s: "AAAAAAAAAAA"},[]string{"AAAAAAAAAA"}},
		{"",args{s: "AAAAAAAAAAAAA"},[]string{"AAAAAAAAAA"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRepeatedDnaSequences(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRepeatedDnaSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
