package leetcode

import "testing"

func Test_getMaximumGenerated(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{n: 7}, 3},
		{"", args{n: 2}, 1},
		{"", args{n: 3}, 2},
		{"", args{n: 0}, 0},
		{"", args{n: 100}, 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMaximumGenerated(tt.args.n); got != tt.want {
				t.Errorf("getMaximumGenerated() = %v, want %v", got, tt.want)
			}
		})
	}
}
