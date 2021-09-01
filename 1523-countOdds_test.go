package leetcode

import "testing"

func Test_countOdds(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{low: 1, high: 2}, 1},
		{"", args{low: 10, high: 20}, 5},
		{"", args{low: 0, high: 4}, 2},
		{"", args{low: 8, high: 10}, 1},
		{"", args{low: 3, high: 7}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOdds(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
