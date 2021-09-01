package leetcode

import "testing"

func Test_reverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"", args{s: "abcdefg", k: 2}, "bacdfeg"},
		{"", args{s: "abcd", k: 2}, "bacd"},
		{"", args{s: "abcde", k: 2}, "bacde"},
		{"", args{s: "abcdef", k: 2}, "bacdfe"},
		{"", args{s: "abcdefg", k: 2}, "bacdfeg"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("reverseStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
