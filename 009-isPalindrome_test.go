package leetcode

import "testing"

func Test_isPalindrome(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"", args{x: 121}, true},
		{"", args{x: 1221}, true},
		{"", args{x: -1}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
