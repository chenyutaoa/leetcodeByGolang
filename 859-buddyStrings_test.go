package leetcode

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{
			s:    "ab",
			goal: "ba",
		}, true},
		{"", args{
			s:    "ab",
			goal: "ab",
		}, false},
		{"", args{
			s:    "aa",
			goal: "aa",
		}, true},
		{"", args{
			s:    "aaaaaaabc",
			goal: "aaaaaaacb",
		}, true},
		{"", args{
			s:    "aba",
			goal: "aab",
		}, true},
		{"", args{
			s:    "aba",
			goal: "cbc",
		}, false},
		{"", args{
			s:    "aba",
			goal: "abc",
		}, false},
		{"", args{
			s:    "abab",
			goal: "abab",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
