package leetcode

import "testing"

func Test_getSum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"",args{
			a: 1,
			b: 2,
		},3},
		{"",args{
			a: -1,
			b: 2,
		},1},
		{"",args{
			a: 257,
			b: -257,
		},0},
		{"",args{
			a: 999,
			b: -257,
		},742},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSum(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
