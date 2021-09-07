package leetcode

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"",args{s: "a"},1},
		{"",args{s: "abcabcbb"},3},
		{"",args{s: "bbbbb"},1},
		{"",args{s: "pwwkew"},3},
		{"",args{s: ""},0},
		{"",args{s: "aab"},2},
		{"",args{s: "cdd"},2},
		{"",args{s: "ab"},2},
		{"",args{s: "dvdf"},3},
		{"",args{s: "abba"},2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
