package leetcode

import "testing"

func Test_reverse(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{x: 1}, 1},
		{"", args{x: 123}, 321},
		{"", args{x: -123}, -321},
		{"", args{x: -100}, -1},
		{"", args{x: 100}, 1},
		{"", args{x: 1534236469}, 0},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
