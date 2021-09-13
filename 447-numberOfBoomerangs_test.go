package leetcode

import "testing"

func Test_numberOfBoomerangs(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{points: [][]int{{0, 0}, {0, 1}, {0, 2}}}, 2},
		{"", args{points: [][]int{{0, 0}, {1, 1}, {2, 2}}}, 2},
		{"", args{points: [][]int{{0, 0}, {1, 0}, {-1,0},{0,1},{0,-1}}}, 20},
		{"", args{points: [][]int{{0, 0}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numberOfBoomerangs(tt.args.points); got != tt.want {
				t.Errorf("numberOfBoomerangs() = %v, want %v", got, tt.want)
			}
		})
	}
}
