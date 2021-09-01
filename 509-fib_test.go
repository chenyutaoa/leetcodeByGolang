package leetcode

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{n: 2}, 1},
		{"", args{n: 3}, 2},
		{"", args{n: 4}, 3},
		{"", args{n: 0}, 0},
		{"", args{n: 1}, 1},
		{"", args{n: 14}, 377},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
